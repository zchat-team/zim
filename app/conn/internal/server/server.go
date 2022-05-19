package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/zmicro-team/zim/pkg/runtime"
	"github.com/zmicro-team/zim/proto/common"
	"github.com/zmicro-team/zim/proto/conn"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/iobrother/ztimer"
	"github.com/nats-io/nats.go"
	"github.com/panjf2000/gnet/pool/goroutine"
	"github.com/zmicro-team/zim/app/conn/internal/client"
	"github.com/zmicro-team/zim/app/conn/protocol"
	"github.com/zmicro-team/zim/proto/sess"
	zerrors "github.com/zmicro-team/zmicro/core/errors"
	"github.com/zmicro-team/zmicro/core/log"
)

const (
	WsUpgrading = 0
	AuthPending = 1
	Authed      = 2
)

type CmdFunc func(client *Client, p *protocol.Packet) (err error)

type Server struct {
	opts      Options
	tcpServer *TcpServer
	wsServer  *WsServer
	timer     *ztimer.Timer
	// TODO: 分桶
	clientManager *ClientManager
	workerPool    *goroutine.Pool
	mapCmdFunc    map[conn.CmdId]CmdFunc
}

type Registry struct {
	BasePath string
	EtcdAddr []string
}

func NewServer(opts ...Option) *Server {
	s := new(Server)
	s.opts = NewOptions(opts...)
	s.clientManager = NewClientManager()
	s.workerPool = goroutine.Default()

	if s.opts.TcpAddr != "" {
		s.tcpServer = NewTcpServer(s, s.opts.TcpAddr)
	}

	if s.opts.WsAddr != "" {
		s.wsServer = NewWsServer(s, s.opts.WsAddr)
	}

	s.timer = ztimer.NewTimer(100*time.Millisecond, 20)

	runtime.Setup()
	s.registerCmdFunc()

	return s
}

func (s *Server) registerCmdFunc() {
	s.mapCmdFunc = make(map[conn.CmdId]CmdFunc)
	s.mapCmdFunc[conn.CmdId_Cmd_Noop] = s.handleNoop
	s.mapCmdFunc[conn.CmdId_Cmd_Logout] = s.handleLogout
	s.mapCmdFunc[conn.CmdId_Cmd_Send] = s.handleSend
	s.mapCmdFunc[conn.CmdId_Cmd_Sync] = s.handleSync
	s.mapCmdFunc[conn.CmdId_Cmd_MsgAck] = s.handleMsgAck

	s.mapCmdFunc[conn.CmdId_Cmd_GetRecentConversation] = s.handleGetRecentConversation
	s.mapCmdFunc[conn.CmdId_Cmd_GetConversationMsg] = s.handleGetConversationMsg
	s.mapCmdFunc[conn.CmdId_Cmd_DeleteConversation] = s.handleDeleteConversation
	s.mapCmdFunc[conn.CmdId_Cmd_GetConversation] = s.handleGetConversation
	s.mapCmdFunc[conn.CmdId_Cmd_SetConversationTop] = s.handleSetConversationTop
	s.mapCmdFunc[conn.CmdId_Cmd_SetConversationMute] = s.handleSetConversationMute
	s.mapCmdFunc[conn.CmdId_Cmd_SyncConversation] = s.handleSyncConversation
	s.mapCmdFunc[conn.CmdId_Cmd_SyncConversationMsg] = s.handleSyncConversationMsg
	s.mapCmdFunc[conn.CmdId_Cmd_SetConversationRead] = s.handleSetConversationRead
}

func (s *Server) GetClientManager() *ClientManager {
	return s.clientManager
}

func (s *Server) GetServerId() string {
	return s.opts.Id
}

func (s *Server) GetTcpServer() *TcpServer {
	return s.tcpServer
}

func (s *Server) GetWsServer() *WsServer {
	return s.wsServer
}

func (s *Server) GetTimer() *ztimer.Timer {
	return s.timer
}

func (s *Server) Start() error {
	go func() {
		if err := s.consumePush(); err != nil {
			log.Error(err)
		}
	}()
	go func() {
		s.timer.Start()
	}()
	go func() {
		if s.tcpServer != nil {
			if err := s.tcpServer.Start(); err != nil {
				log.Error(err)
			}
		}
	}()
	go func() {
		if s.wsServer != nil {
			if err := s.wsServer.Start(); err != nil {
				log.Error(err)
			}
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	var lastError error
	if s.tcpServer != nil {
		if err := s.tcpServer.Stop(); err != nil {
			lastError = err
		}
	}
	if s.wsServer != nil {
		if err := s.wsServer.Stop(); err != nil {
			lastError = err
		}
	}
	return lastError
}

func (s *Server) consumeKick() error {
	return nil
}

func (s *Server) consumePush() error {
	// process push message
	pushMsg := new(common.PushMsg)
	topic := fmt.Sprintf("push.online.%s", s.GetServerId())
	nc := runtime.GetNC()
	if _, err := nc.Subscribe(topic, func(msg *nats.Msg) {

		if err := proto.Unmarshal(msg.Data, pushMsg); err != nil {
			log.Errorf("proto.Unmarshal error=(%v)", err)
			return
		}

		log.Infof("recv a msg=%v", pushMsg)
		for _, deviceId := range pushMsg.Devices {
			if c := s.GetClientManager().Get(deviceId); c != nil {
				if c.Conn != nil {
					p := protocol.Packet{
						HeaderLen: 20,
						Version:   uint32(c.Version),
						Cmd:       uint32(conn.CmdId_Cmd_Msg),
						Seq:       0,
						BodyLen:   uint32(len(pushMsg.Msg)),
						Body:      pushMsg.Msg,
					}
					c.WritePacket(&p)
				}
			}
		}
	}); err != nil {
		return err
	}
	return nil
}

func (s *Server) OnOpen(client *Client) {
	// 10秒钟之内没有认证成功，关闭连接
	client.TimerTask = s.GetTimer().AfterFunc(time.Second*10, func() {
		log.Info("auth timeout...")
		client.Close()
	})
}

func (s *Server) OnClose(c *Client) {
	log.Infof("client=%s close", c.Uin)

	if c.DeviceId != "" {
		s.GetClientManager().Remove(c)
	}

	if c.Status != Authed {
		return
	}

	s.workerPool.Submit(func() {
		if c != nil {
			req := sess.LogoutReq{
				Uin:      c.Uin,
				DeviceId: c.DeviceId,
			}
			client.GetSessClient().Logout(context.Background(), &req)
		}
	})
}

func (s *Server) OnMessage(data []byte, client *Client) {
	s.workerPool.Submit(func() {
		p := &protocol.Packet{}
		if err := p.Read(data); err != nil {
			log.Error(err)
			client.Close()
			return
		}

		if client.Status == AuthPending {
			cmd := conn.CmdId(p.Cmd)
			if cmd != conn.CmdId_Cmd_Login {
				log.Error("first packet must be cmd_login")
				client.Close()
				return
			}
			if err := s.handleLogin(client, p); err != nil {
				client.Close()
				log.Info("AUTH FAILED")
			} else {
				client.Status = Authed
			}
		} else {
			s.handleProto(client, p)
		}
	})

}

func (s *Server) handleLogin(c *Client, p *protocol.Packet) (err error) {
	req := &conn.LoginReq{}

	rsp := &conn.LoginRsp{
		Code:    200,
		Message: "成功",
	}

	defer func() {
		// 不论登录成功与失败，均取消定时任务
		c.TimerTask.Cancel()
		c.TimerTask = nil

		b, err := proto.Marshal(rsp)
		if err != nil {
			rsp.Code = 500
			rsp.Message = err.Error()
			log.Error(err)
		}

		p.BodyLen = uint32(len(b))
		p.Body = b
		if err := c.WritePacket(p); err != nil {
			log.Error(err)
		}
	}()

	if err = proto.Unmarshal(p.Body, req); err != nil {
		log.Error(err)
		rsp.Code = 500
		rsp.Message = "协议解析错误"
		err = errors.New("协议解析错误")
		return
	}

	if req.Uin == "" {
		rsp.Code = 500
		rsp.Message = "账号不能为空"
		log.Error("账号不能为空")
		err = errors.New("账号不能为空")
		return
	}

	// TODO: DeviceId -> ConnId，服务端来生成
	log.Infof("handleLogin uin=%s platform=%s token=%s device_id=%s device_name=%s",
		req.Uin, req.Platform, req.Token, req.DeviceId, req.DeviceName)

	reqL := sess.LoginReq{
		Uin:        req.Uin,
		Platform:   req.Platform,
		Server:     s.GetServerId(),
		Token:      req.Token,
		DeviceId:   req.DeviceId,
		DeviceName: req.DeviceName,
		Tag:        req.Tag,
		Reconnect:  req.Reconnect,
	}
	rspL, err := client.GetSessClient().Login(context.Background(), &reqL)
	if err != nil {
		// TODO:
		e := zerrors.FromError(err)
		rsp.Code = e.Code
		rsp.Message = e.Message
		if e.Message == "" {
			rsp.Message = e.Detail
		}
		return
	}

	rsp.Code = rspL.Code
	rsp.Message = rspL.Message

	if req.Reconnect && rsp.Code == 409 {
		log.Infof("登录冲突 uin=%s cur_device_id=%s conflict_device_id=%s conflict_device_name=%s",
			req.Uin, req.DeviceId, rspL.ConflictDeviceId, rspL.ConflictDeviceName)
		return
	}
	// 踢掉旧的连接
	if rspL.ConflictDeviceId != "" {
		log.Infof("conflict device id=%s", rspL.ConflictDeviceId)
		oldClient := s.GetClientManager().Get(rspL.ConflictDeviceId)
		if oldClient != nil && oldClient.Conn != nil {
			reason := fmt.Sprintf("您的账号在设备%s上登录，如果不是本人操作，您的账号可能被盗", req.DeviceName)
			kick := &conn.Kick{Reason: reason}
			if b, err := proto.Marshal(kick); err == nil {
				pp := protocol.Packet{
					HeaderLen: p.HeaderLen,
					Version:   p.Version,
					Cmd:       uint32(conn.CmdId_Cmd_Kick),
					Seq:       0,
					BodyLen:   uint32(len(b)),
					Body:      b,
				}

				oldClient.WritePacket(&pp)
			}
			log.Infof("踢掉客户端 uin=%s device_id=%s", oldClient.Uin, oldClient.DeviceId)
			oldClient.Close()
			s.GetClientManager().Remove(oldClient)
		}
	}

	c.DeviceId = req.DeviceId
	c.Uin = reqL.Uin
	c.Platform = req.Platform
	c.Server = s.GetServerId()
	c.Version = int(p.Version)
	s.GetClientManager().Add(c)

	log.Infof("AUTH SUCC uin=%s", req.Uin)
	return
}

func (s *Server) handleLogout(c *Client, p *protocol.Packet) (err error) {
	log.Infof("client %s noop", c.Uin)
	c.WritePacket(p)
	req := sess.LogoutReq{
		Uin:      c.Uin,
		DeviceId: c.DeviceId,
	}
	client.GetSessClient().Logout(context.Background(), &req)

	return
}

func (s *Server) handleProto(c *Client, p *protocol.Packet) (err error) {
	log.Infof("cmd=%d", p.Cmd)
	cmd := conn.CmdId(p.Cmd)

	if s.mapCmdFunc[cmd] != nil {
		err = s.mapCmdFunc[cmd](c, p)
	}

	return
}

func (s *Server) handleNoop(c *Client, p *protocol.Packet) (err error) {
	log.Infof("client %s noop", c.Uin)
	c.WritePacket(p)
	req := sess.HeartbeatReq{
		Uin:      c.Uin,
		DeviceId: c.DeviceId,
		Server:   c.Server,
	}
	client.GetSessClient().Heartbeat(context.Background(), &req)

	return
}
