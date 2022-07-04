package service

import (
	"context"
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
	var devices []*DeviceInfo
	if req.Tag != "" {
		v, e := s.getOnlineOfTag(ctx, req.Uin, req.Tag)
		if e != nil {
			err = e
			return
		}
		devices = append(devices, v...)
	}

	if len(devices) > 0 {
		if req.Reconnect {
			rsp.ConflictDeviceId = devices[len(devices)-1].DeviceId
			rsp.ConflictDeviceName = devices[len(devices)-1].DeviceName
			return
		}

		rsp.ConflictDeviceId = devices[0].DeviceId
		rsp.ConflictDeviceName = devices[0].DeviceName
		devices[0].DisconnectTime = time.Now().Unix()
		s.delConn(ctx, req.Uin, devices[0])
	}

	info := &DeviceInfo{
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

	return
}

func (s *Service) Logout(ctx context.Context, req *sess.LogoutReq, rsp *sess.LogoutRsp) (err error) {
	info := s.getDevice(ctx, req.Uin, req.DeviceId)
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
	info := s.getDevice(ctx, req.Uin, req.DeviceId)
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
	info := s.getDevice(ctx, req.Uin, req.DeviceId)
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
	receivers, _ := s.getOnline(ctx, req.Uin)
	for server, devices := range receivers {
		for _, d := range devices {
			item := sess.DeviceInfo{
				DeviceId: d.DeviceId,
				Server:   server,
				Status:   int32(d.GetRealStatus()),
			}
			rsp.Devices = append(rsp.Devices, &item)
		}
	}
	return
}
