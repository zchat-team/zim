package chat

import (
	"context"
	"sync"

	zhttp "github.com/zmicro-team/zmicro/core/transport/http"

	"github.com/zchat-team/zim/app/rest/internal/client"
	"github.com/zchat-team/zim/proto/http/rest/chat"
	pb "github.com/zchat-team/zim/proto/rpc/chat"
)

type Service struct {
	zhttp.Implemented
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

func (s *Service) Send(ctx context.Context, req *chat.SendReq, rsp *chat.SendRsp) (err error) {
	reqL := pb.SendReq{
		ConvType:      req.ConvType,
		MsgType:       req.MsgType,
		Sender:        req.Sender,
		Target:        req.Target,
		Content:       req.Content,
		AtUserList:    nil,
		ClientUuid:    "",
		IsTransparent: req.IsTransparent,
	}

	cli := client.GetChatClient()
	rspL, err := cli.SendMsg(ctx, &reqL)
	if err != nil {
		return
	}

	rsp.Id = rspL.Id
	rsp.SendTime = rspL.SendTime

	return
}
