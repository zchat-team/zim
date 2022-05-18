package server

import (
	"github.com/google/uuid"
	"github.com/panjf2000/gnet/pool/goroutine"
	"github.com/zmicro-team/zim/app/conn/protocol"
	"github.com/zmicro-team/zim/proto/conn"
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

	s.opts.Id = uuid.New().String()

	if s.opts.TcpAddr != "" {
		s.tcpServer = NewTcpServer(s, s.opts.TcpAddr)
	}

	if s.opts.WsAddr != "" {
		s.wsServer = NewWsServer(s, s.opts.WsAddr)
	}

	return s
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

func (s *Server) Start() error {

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

func (s *Server) OnOpen(client *Client) {
	log.Info("client open")
}

func (s *Server) OnClose(c *Client) {
	log.Infof("client=%s close", c.Uin)

	if c.DeviceId != "" {
		s.GetClientManager().Remove(c)
	}
}

func (s *Server) OnMessage(data []byte, client *Client) {
	s.workerPool.Submit(func() {
		p := &protocol.Packet{}
		if err := p.Read(data); err != nil {
			log.Error(err)
			client.Close()
			return
		}

		// TODO: handle
	})
}
