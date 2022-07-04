package service

import (
	"context"
	"github.com/google/uuid"
	"sync"
	"time"

	"github.com/zchat-team/zim/pkg/constant"
	"github.com/zchat-team/zim/proto/rpc/sess"
)

type Service struct {
}

var (
	service *Service
	once    sync.Once
)

func GetService() *Service {
	once.Do(func() {
		service = &Service{}
	})
	return service
}

func (s *Service) Login(ctx context.Context, req *sess.LoginReq, rsp *sess.LoginRsp) (err error) {
	// TODO: token验证
	var onlines []*ConnInfo
	if req.Tag != "" {
		v, e := s.getOnlineOfTag(ctx, req.Uin, req.Tag)
		if e != nil {
			err = e
			return
		}
		onlines = append(onlines, v...)
	}

	if len(onlines) > 0 {
		if req.Reconnect {
			rsp.ConflictDeviceId = onlines[len(onlines)-1].DeviceId
			rsp.ConflictDeviceName = onlines[len(onlines)-1].DeviceName
			return
		}

		rsp.ConflictDeviceId = onlines[0].DeviceId
		rsp.ConflictDeviceName = onlines[0].DeviceName
		onlines[0].DisconnectTime = time.Now().Unix()
		s.delConn(ctx, req.Uin, onlines[0])
	}

	info := &ConnInfo{
		ConnID:         uuid.New().String(),
		DeviceId:       req.DeviceId,
		DeviceName:     req.DeviceName,
		Tag:            req.Tag,
		Platform:       req.Platform,
		Server:         req.Server,
		LoginTime:      time.Now().Unix(),
		DisconnectTime: 0,
		Status:         constant.Online,
	}
	if err = s.addConn(ctx, req.Uin, info); err != nil {
		return
	}

	rsp.ConnId = info.ConnID

	return
}

func (s *Service) Logout(ctx context.Context, req *sess.LogoutReq, rsp *sess.LogoutRsp) (err error) {
	info := s.getConn(ctx, req.Uin, req.ConnId)
	if info == nil {
		return
	}

	info.DisconnectTime = time.Now().Unix()
	info.Status = constant.Offline

	if err = s.delConn(ctx, req.Uin, info); err != nil {
		return
	}
	*rsp = sess.LogoutRsp{}
	return
}

func (s *Service) Disconnect(ctx context.Context, req *sess.DisconnectReq, rsp *sess.DisconnectRsp) (err error) {
	info := s.getConn(ctx, req.Uin, req.ConnId)
	if info == nil {
		return
	}

	info.DisconnectTime = time.Now().Unix()
	info.Status = constant.PushOnline

	if err = s.delConn(ctx, req.Uin, info); err != nil {
		return
	}
	*rsp = sess.DisconnectRsp{}
	return
}

func (s *Service) Heartbeat(ctx context.Context, req *sess.HeartbeatReq, rsp *sess.HeartbeatRsp) (err error) {
	info := s.getConn(ctx, req.Uin, req.ConnId)
	if info == nil {
		return
	}
	info.DisconnectTime = 0
	info.Status = constant.Online

	if err = s.addConn(ctx, req.Uin, info); err != nil {
		return
	}

	*rsp = sess.HeartbeatRsp{}
	return
}

func (s *Service) GetOnline(ctx context.Context, req *sess.GetOnlineReq, rsp *sess.GetOnlineRsp) (err error) {
	onlines, _ := s.getOnline(ctx, req.Uin)
	for server, conns := range onlines {
		for _, d := range conns {
			item := sess.ConnInfo{
				ConnId:   d.ConnID,
				DeviceId: d.DeviceId,
				Server:   server,
				Status:   int32(d.GetRealStatus()),
			}
			rsp.Conns = append(rsp.Conns, &item)
		}
	}
	return
}
