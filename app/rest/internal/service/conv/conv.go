package conv

import (
	"context"
	"github.com/zchat-team/zim/api/rest/conv"
	zgin "github.com/zmicro-team/zmicro/core/transport/http"
	"sync"
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

func (s *Service) AddMember(ctx context.Context, req *conv.AddMemberReq, rsp *conv.AddMemberRsp) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) AddMuteClient(ctx context.Context, req *conv.AddMuteClientReq, rsp *conv.AddMuteClientRsp) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Create(ctx context.Context, req *conv.CreateReq, rsp *conv.CreateRsp) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Delete(ctx context.Context, req *conv.DeleteReq, rsp *conv.DeleteRsp) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) DeleteMember(ctx context.Context, req *conv.DeleteMemberReq, rsp *conv.DeleteMemberRsp) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) DeleteMsg(ctx context.Context, req *conv.DeleteMsgReq, rsp *conv.DeleteMsgRsp) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) DeleteMuteClient(ctx context.Context, req *conv.DeleteMuteClientReq, rsp *conv.DeleteMuteClientRsp) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Query(ctx context.Context, req *conv.QueryReq, rsp *conv.QueryRsp) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) QueryMember(ctx context.Context, req *conv.QueryMemberReq, rsp *conv.QueryMemberRsp) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) QueryMsg(ctx context.Context, req *conv.QueryMsgReq, rsp *conv.QueryMsgMsgRsp) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) QueryMuteClient(ctx context.Context, req *conv.QueryMuteClientReq, rsp *conv.QueryMuteClientRsp) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Recall(ctx context.Context, req *conv.RecallReq, rsp *conv.RecallRsp) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Send(ctx context.Context, req *conv.SendReq, rsp *conv.SendRsp) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Update(ctx context.Context, req *conv.UpdateReq, rsp *conv.UpdateRsp) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateMsg(ctx context.Context, req *conv.UpdateMsgReq, rsp *conv.UpdateMsgRsp) (err error) {
	//TODO implement me
	panic("implement me")
}
