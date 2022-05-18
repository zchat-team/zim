package service

import (
	"context"
	"github.com/zmicro-team/zim/proto/chat"
	"github.com/zmicro-team/zim/proto/common"
)

type Chat struct {
}

func GetChatService() *Chat {
	c := &Chat{}
	return c
}

func (l *Chat) SendMsg(ctx context.Context, req *chat.SendReq, rsp *chat.SendRsp) (err error) {
	return
}

func (l *Chat) SyncMsg(ctx context.Context, req *chat.SyncMsgReq, rsp *chat.SyncMsgRsp) (err error) {

	return
}

func (l *Chat) removeDirty(ctx context.Context, req *chat.SyncMsgReq) (err error) {
	return
}

func (l *Chat) MsgAck(ctx context.Context, req *chat.MsgAckReq, rsp *chat.MsgAckRsp) (err error) {

	return
}

func (l *Chat) sendC2C(ctx context.Context, req *chat.SendReq, rsp *chat.SendRsp) (err error) {
	return
}

func (l *Chat) sendC2G(ctx context.Context, req *chat.SendReq, rsp *chat.SendRsp) (err error) {

	return
}

func (l *Chat) createConversation(ctx context.Context, owner, target string, convType int, m *common.Msg) (err error) {

	return
}

func (l *Chat) Recall(ctx context.Context, req *chat.RecallReq, rsp *chat.RecallRsp) (err error) {
	return
}

func (l *Chat) DeleteMsg(ctx context.Context, req *chat.DeleteMsgReq, rsp *chat.DeleteMsgRsp) (err error) {

	return
}
