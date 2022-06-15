package group

import (
	"context"
	"github.com/zchat-team/zim/app/rest/internal/client"
	"github.com/zchat-team/zim/proto/http/rest/group"
	pb "github.com/zchat-team/zim/proto/rpc/group"

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

func (s *Service) Create(ctx context.Context, req *group.CreateReq, rsp *group.CreateRsp) (err error) {

	reqL := pb.CreateReq{
		Owner:   req.Owner,
		Members: req.Members,
		Name:    req.Name,
		GroupId: req.GroupId,
		Notice:  req.Notice,
		Intro:   req.Intro,
		Avatar:  req.Avatar,
	}

	cli := client.GetGroupClient()
	rspL, err := cli.Create(ctx, &reqL)
	if err != nil {
		return
	}
	rsp.GroupId = rspL.GroupId
	return
}

func (s *Service) Add(ctx context.Context, req *group.AddReq, rsp *group.AddRsp) (err error) {
	reqL := pb.InviteUserToGroupReq{
		GroupId:  req.GroupId,
		UserList: req.Members,
	}

	cli := client.GetGroupClient()
	_, err = cli.InviteUserToGroup(ctx, &reqL)

	return
}
