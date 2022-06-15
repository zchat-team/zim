// Code generated by protoc-gen-rpcx. DO NOT EDIT.
// versions:
// - protoc-gen-rpcx v0.3.0
// - protoc          v3.19.0
// source: proto/rpc/chat/conv.proto

package chat

import (
	context "context"
	client "github.com/smallnest/rpcx/client"
	protocol "github.com/smallnest/rpcx/protocol"
	server "github.com/smallnest/rpcx/server"
	common "github.com/zchat-team/zim/proto/rpc/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = context.TODO
var _ = server.NewServer
var _ = client.NewClient
var _ = protocol.NewMessage

//================== interface skeleton ===================
type ConvAble interface {
	// ConvAble can be used for interface verification.

	// GetRecentConversation is server rpc method as defined
	GetRecentConversation(ctx context.Context, args *GetRecentConversationReq, reply *GetRecentConversationRsp) (err error)

	// GetConversationMsg is server rpc method as defined
	GetConversationMsg(ctx context.Context, args *GetConversationMsgReq, reply *GetConversationMsgRsp) (err error)

	// DeleteConversation is server rpc method as defined
	DeleteConversation(ctx context.Context, args *DeleteConversationReq, reply *DeleteConversationRsp) (err error)

	// GetConversation is server rpc method as defined
	GetConversation(ctx context.Context, args *GetConversationReq, reply *common.Conversation) (err error)

	// SetConversationTop is server rpc method as defined
	SetConversationTop(ctx context.Context, args *SetConversationTopReq, reply *SetConversationTopRsp) (err error)

	// SetConversationMute is server rpc method as defined
	SetConversationMute(ctx context.Context, args *SetConversationMuteReq, reply *SetConversationMuteRsp) (err error)

	// SetConversationRead is server rpc method as defined
	SetConversationRead(ctx context.Context, args *SetConversationReadReq, reply *SetConversationReadRsp) (err error)

	// SyncConversation is server rpc method as defined
	SyncConversation(ctx context.Context, args *SyncConversationReq, reply *SyncConversationRsp) (err error)

	// SyncConversationMsg is server rpc method as defined
	SyncConversationMsg(ctx context.Context, args *SyncConversationMsgReq, reply *SyncConversationMsgRsp) (err error)
}

//================== server skeleton ===================
type ConvImpl struct{}

// ServeForConv starts a server only registers one service.
// You can register more services and only start one server.
// It blocks until the application exits.
func ServeForConv(addr string) error {
	s := server.NewServer()
	s.RegisterName("Conv", new(ConvImpl), "")
	return s.Serve("tcp", addr)
}

// GetRecentConversation is server rpc method as defined
func (s *ConvImpl) GetRecentConversation(ctx context.Context, args *GetRecentConversationReq, reply *GetRecentConversationRsp) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = GetRecentConversationRsp{}

	return nil
}

// GetConversationMsg is server rpc method as defined
func (s *ConvImpl) GetConversationMsg(ctx context.Context, args *GetConversationMsgReq, reply *GetConversationMsgRsp) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = GetConversationMsgRsp{}

	return nil
}

// DeleteConversation is server rpc method as defined
func (s *ConvImpl) DeleteConversation(ctx context.Context, args *DeleteConversationReq, reply *DeleteConversationRsp) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = DeleteConversationRsp{}

	return nil
}

// GetConversation is server rpc method as defined
func (s *ConvImpl) GetConversation(ctx context.Context, args *GetConversationReq, reply *common.Conversation) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = common.Conversation{}

	return nil
}

// SetConversationTop is server rpc method as defined
func (s *ConvImpl) SetConversationTop(ctx context.Context, args *SetConversationTopReq, reply *SetConversationTopRsp) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = SetConversationTopRsp{}

	return nil
}

// SetConversationMute is server rpc method as defined
func (s *ConvImpl) SetConversationMute(ctx context.Context, args *SetConversationMuteReq, reply *SetConversationMuteRsp) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = SetConversationMuteRsp{}

	return nil
}

// SetConversationRead is server rpc method as defined
func (s *ConvImpl) SetConversationRead(ctx context.Context, args *SetConversationReadReq, reply *SetConversationReadRsp) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = SetConversationReadRsp{}

	return nil
}

// SyncConversation is server rpc method as defined
func (s *ConvImpl) SyncConversation(ctx context.Context, args *SyncConversationReq, reply *SyncConversationRsp) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = SyncConversationRsp{}

	return nil
}

// SyncConversationMsg is server rpc method as defined
func (s *ConvImpl) SyncConversationMsg(ctx context.Context, args *SyncConversationMsgReq, reply *SyncConversationMsgRsp) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = SyncConversationMsgRsp{}

	return nil
}

//================== client stub ===================
// Conv is a client wrapped XClient.
type ConvClient struct {
	xclient client.XClient
}

// NewConvClient wraps a XClient as ConvClient.
// You can pass a shared XClient object created by NewXClientForConv.
func NewConvClient(xclient client.XClient) *ConvClient {
	return &ConvClient{xclient: xclient}
}

// NewXClientForConv creates a XClient.
// You can configure this client with more options such as etcd registry, serialize type, select algorithm and fail mode.
func NewXClientForConv(addr string) (client.XClient, error) {
	d, err := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	if err != nil {
		return nil, err
	}

	opt := client.DefaultOption
	opt.SerializeType = protocol.ProtoBuffer

	xclient := client.NewXClient("Conv", client.Failtry, client.RoundRobin, d, opt)

	return xclient, nil
}

// GetRecentConversation is client rpc method as defined
func (c *ConvClient) GetRecentConversation(ctx context.Context, args *GetRecentConversationReq) (reply *GetRecentConversationRsp, err error) {
	reply = &GetRecentConversationRsp{}
	err = c.xclient.Call(ctx, "GetRecentConversation", args, reply)
	return reply, err
}

// GetConversationMsg is client rpc method as defined
func (c *ConvClient) GetConversationMsg(ctx context.Context, args *GetConversationMsgReq) (reply *GetConversationMsgRsp, err error) {
	reply = &GetConversationMsgRsp{}
	err = c.xclient.Call(ctx, "GetConversationMsg", args, reply)
	return reply, err
}

// DeleteConversation is client rpc method as defined
func (c *ConvClient) DeleteConversation(ctx context.Context, args *DeleteConversationReq) (reply *DeleteConversationRsp, err error) {
	reply = &DeleteConversationRsp{}
	err = c.xclient.Call(ctx, "DeleteConversation", args, reply)
	return reply, err
}

// GetConversation is client rpc method as defined
func (c *ConvClient) GetConversation(ctx context.Context, args *GetConversationReq) (reply *common.Conversation, err error) {
	reply = &common.Conversation{}
	err = c.xclient.Call(ctx, "GetConversation", args, reply)
	return reply, err
}

// SetConversationTop is client rpc method as defined
func (c *ConvClient) SetConversationTop(ctx context.Context, args *SetConversationTopReq) (reply *SetConversationTopRsp, err error) {
	reply = &SetConversationTopRsp{}
	err = c.xclient.Call(ctx, "SetConversationTop", args, reply)
	return reply, err
}

// SetConversationMute is client rpc method as defined
func (c *ConvClient) SetConversationMute(ctx context.Context, args *SetConversationMuteReq) (reply *SetConversationMuteRsp, err error) {
	reply = &SetConversationMuteRsp{}
	err = c.xclient.Call(ctx, "SetConversationMute", args, reply)
	return reply, err
}

// SetConversationRead is client rpc method as defined
func (c *ConvClient) SetConversationRead(ctx context.Context, args *SetConversationReadReq) (reply *SetConversationReadRsp, err error) {
	reply = &SetConversationReadRsp{}
	err = c.xclient.Call(ctx, "SetConversationRead", args, reply)
	return reply, err
}

// SyncConversation is client rpc method as defined
func (c *ConvClient) SyncConversation(ctx context.Context, args *SyncConversationReq) (reply *SyncConversationRsp, err error) {
	reply = &SyncConversationRsp{}
	err = c.xclient.Call(ctx, "SyncConversation", args, reply)
	return reply, err
}

// SyncConversationMsg is client rpc method as defined
func (c *ConvClient) SyncConversationMsg(ctx context.Context, args *SyncConversationMsgReq) (reply *SyncConversationMsgRsp, err error) {
	reply = &SyncConversationMsgRsp{}
	err = c.xclient.Call(ctx, "SyncConversationMsg", args, reply)
	return reply, err
}

//================== oneclient stub ===================
// ConvOneClient is a client wrapped oneClient.
type ConvOneClient struct {
	serviceName string
	oneclient   *client.OneClient
}

// NewConvOneClient wraps a OneClient as ConvOneClient.
// You can pass a shared OneClient object created by NewOneClientForConv.
func NewConvOneClient(oneclient *client.OneClient) *ConvOneClient {
	return &ConvOneClient{
		serviceName: "Conv",
		oneclient:   oneclient,
	}
}

// ======================================================

// GetRecentConversation is client rpc method as defined
func (c *ConvOneClient) GetRecentConversation(ctx context.Context, args *GetRecentConversationReq) (reply *GetRecentConversationRsp, err error) {
	reply = &GetRecentConversationRsp{}
	err = c.oneclient.Call(ctx, c.serviceName, "GetRecentConversation", args, reply)
	return reply, err
}

// GetConversationMsg is client rpc method as defined
func (c *ConvOneClient) GetConversationMsg(ctx context.Context, args *GetConversationMsgReq) (reply *GetConversationMsgRsp, err error) {
	reply = &GetConversationMsgRsp{}
	err = c.oneclient.Call(ctx, c.serviceName, "GetConversationMsg", args, reply)
	return reply, err
}

// DeleteConversation is client rpc method as defined
func (c *ConvOneClient) DeleteConversation(ctx context.Context, args *DeleteConversationReq) (reply *DeleteConversationRsp, err error) {
	reply = &DeleteConversationRsp{}
	err = c.oneclient.Call(ctx, c.serviceName, "DeleteConversation", args, reply)
	return reply, err
}

// GetConversation is client rpc method as defined
func (c *ConvOneClient) GetConversation(ctx context.Context, args *GetConversationReq) (reply *common.Conversation, err error) {
	reply = &common.Conversation{}
	err = c.oneclient.Call(ctx, c.serviceName, "GetConversation", args, reply)
	return reply, err
}

// SetConversationTop is client rpc method as defined
func (c *ConvOneClient) SetConversationTop(ctx context.Context, args *SetConversationTopReq) (reply *SetConversationTopRsp, err error) {
	reply = &SetConversationTopRsp{}
	err = c.oneclient.Call(ctx, c.serviceName, "SetConversationTop", args, reply)
	return reply, err
}

// SetConversationMute is client rpc method as defined
func (c *ConvOneClient) SetConversationMute(ctx context.Context, args *SetConversationMuteReq) (reply *SetConversationMuteRsp, err error) {
	reply = &SetConversationMuteRsp{}
	err = c.oneclient.Call(ctx, c.serviceName, "SetConversationMute", args, reply)
	return reply, err
}

// SetConversationRead is client rpc method as defined
func (c *ConvOneClient) SetConversationRead(ctx context.Context, args *SetConversationReadReq) (reply *SetConversationReadRsp, err error) {
	reply = &SetConversationReadRsp{}
	err = c.oneclient.Call(ctx, c.serviceName, "SetConversationRead", args, reply)
	return reply, err
}

// SyncConversation is client rpc method as defined
func (c *ConvOneClient) SyncConversation(ctx context.Context, args *SyncConversationReq) (reply *SyncConversationRsp, err error) {
	reply = &SyncConversationRsp{}
	err = c.oneclient.Call(ctx, c.serviceName, "SyncConversation", args, reply)
	return reply, err
}

// SyncConversationMsg is client rpc method as defined
func (c *ConvOneClient) SyncConversationMsg(ctx context.Context, args *SyncConversationMsgReq) (reply *SyncConversationMsgRsp, err error) {
	reply = &SyncConversationMsgRsp{}
	err = c.oneclient.Call(ctx, c.serviceName, "SyncConversationMsg", args, reply)
	return reply, err
}