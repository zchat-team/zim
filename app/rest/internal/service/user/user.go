package user

import (
	"context"
	"github.com/zchat-team/zim/proto/http/rest/user"
	"sync"

	zgin "github.com/zmicro-team/zmicro/core/transport/http"
)

type Service struct {
	zgin.Implemented
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

func (s *Service) CheckOnline(ctx context.Context, req *user.CheckOnlineReq, rsp *user.CheckOnlineRsp) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetMessage(ctx context.Context, req *user.GetMessageReq, rsp *user.GetMessageRsp) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Kick(ctx context.Context, req *user.KickReq, rsp *user.KickRsp) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UnreadCount(ctx context.Context, req *user.UnreadCountReq, rsp *user.UnreadCountRsp) error {
	//TODO implement me
	panic("implement me")
}
