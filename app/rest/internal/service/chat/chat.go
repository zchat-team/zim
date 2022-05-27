package im

import (
	"context"
	"sync"

	"github.com/zchat-team/zim/api/rest/chat"
	"github.com/zchat-team/zim/app/rest/internal/client"
	pb "github.com/zchat-team/zim/proto/chat"
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

type Registry struct {
	BasePath string
	EtcdAddr []string
}

func (s *Service) Send(ctx context.Context, req *chat.SendReq, rsp *chat.SendRsp) (err error) {
	reqL := pb.SendReq{
		DeviceId:      "",
		ConvType:      req.ConvType,
		MsgType:       req.MsgType,
		Sender:        req.Sender,
		Target:        req.Target,
		Content:       req.Content,
		Extra:         req.Extra,
		AtUserList:    nil,
		IsTransparent: req.IsTransparent,
		ClientUuid:    "",
	}

	cli := client.GetChatClient()
	rspL, err := cli.SendMsg(ctx, &reqL)
	if err != nil {
		return
	}

	rsp = &chat.SendRsp{
		Id:       rspL.Id,
		SendTime: rspL.SendTime,
	}
	return
}
