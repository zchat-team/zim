package group

import (
	"context"

	"sync"

	"github.com/zchat-team/zim/api/rest/group"
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

func (s *Service) Create(ctx context.Context, req *group.CreateReq, rsp *group.CreateRsp) error {
	//TODO implement me
	panic("implement me")
}
