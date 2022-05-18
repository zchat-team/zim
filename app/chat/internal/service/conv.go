package service

import (
	"context"
	"github.com/zmicro-team/zim/proto/chat"
	"github.com/zmicro-team/zim/proto/common"
)

type Conv struct {
}

func GetConvService() *Chat {
	c := &Chat{}
	return c
}

func (l *Conv) GetRecentConversation(ctx context.Context, req *chat.GetRecentConversationReq, rsp *chat.GetRecentConversationRsp) (err error) {

	return
}

func (l *Conv) GetConversationMsg(ctx context.Context, req *chat.GetConversationMsgReq, rsp *chat.GetConversationMsgRsp) (err error) {
	return
}

func (l *Conv) DeleteConversation(ctx context.Context, req *chat.DeleteConversationReq, rsp *chat.DeleteConversationRsp) (err error) {
	return
}

func (l *Conv) GetConversation(ctx context.Context, req *chat.GetConversationReq, rsp *common.Conversation) (err error) {

	return
}

func (l *Conv) SetConversationTop(ctx context.Context, req *chat.SetConversationTopReq, rsp *chat.SetConversationTopRsp) (err error) {
	return
}

func (l *Conv) SetConversationMute(ctx context.Context, req *chat.SetConversationMuteReq, rsp *chat.SetConversationMuteRsp) (err error) {
	return
}

func (l *Conv) SetConversationRead(ctx context.Context, req *chat.SetConversationReadReq, rsp *chat.SetConversationReadRsp) (err error) {
	return
}

func (l *Conv) SyncConversation(ctx context.Context, req *chat.SyncConversationReq, rsp *chat.SyncConversationRsp) (err error) {
	return
}

func (l *Conv) SyncConversationMsg(ctx context.Context, req *chat.SyncConversationMsgReq, rsp *chat.SyncConversationMsgRsp) (err error) {
	return
}

