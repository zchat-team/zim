package server

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/zmicro-team/zim/app/conn/internal/client"
	"github.com/zmicro-team/zim/app/conn/protocol"
	"github.com/zmicro-team/zim/proto/chat"
	zerrors "github.com/zmicro-team/zmicro/core/errors"
	"github.com/zmicro-team/zmicro/core/log"
)

func (s *Server) handleMsgAck(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.MsgAckReq{}
	rsp := &protocol.MsgAckRsp{}

	defer func() {
		b, err := proto.Marshal(rsp)
		if err != nil {
			return
		}

		p.BodyLen = uint32(len(b))
		p.Body = b
		c.WritePacket(p)
	}()

	if err = proto.Unmarshal(p.Body, req); err != nil {
		return
	}

	reqL := chat.MsgAckReq{
		Uin:      c.Uin,
		DeviceId: c.DeviceId,
		Id:       req.Id,
	}

	_, err = client.GetChatClient().MsgAck(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (s *Server) handleSync(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.SyncMsgReq{}
	rsp := &protocol.SyncMsgRsp{}

	defer func() {
		b, err := proto.Marshal(rsp)
		if err != nil {
			return
		}

		p.BodyLen = uint32(len(b))
		p.Body = b
		c.WritePacket(p)
	}()

	if err = proto.Unmarshal(p.Body, req); err != nil {
		return
	}

	reqL := chat.SyncMsgReq{
		Uin:      c.Uin,
		DeviceId: c.DeviceId,
		Offset:   req.Offset,
		Limit:    req.Limit,
	}

	rspL, err := client.GetChatClient().SyncMsg(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	for _, v := range rspL.List {
		msg := &protocol.Msg{
			Id:         v.Id,
			ConvType:   v.ConvType,
			Type:       v.Type,
			Content:    v.Content,
			Sender:     v.Sender,
			Target:     v.Target,
			Extra:      v.Extra,
			SendTime:   v.SendTime,
			AtUserList: v.AtUserList,
			ReadTime:   v.ReadTime,
			ClientUuid: v.ClientUuid,
		}
		rsp.List = append(rsp.List, msg)
	}

	return
}

func (s *Server) handleSend(c *Client, p *protocol.Packet) (err error) {
	log.Info("handleSend ...")
	req := &protocol.SendReq{}

	rsp := &protocol.SendRsp{
		Code:    200,
		Message: "成功",
	}

	defer func() {
		b, err := proto.Marshal(rsp)
		if err != nil {
			return
		}

		p.BodyLen = uint32(len(b))
		p.Body = b
		c.WritePacket(p)
	}()

	if err = proto.Unmarshal(p.Body, req); err != nil {
		rsp.Code = 500
		rsp.Message = "协议解析错误"
		log.Error(err)
		return
	}

	r := chat.SendReq{
		DeviceId:      c.DeviceId,
		ConvType:      req.ConvType,
		MsgType:       req.MsgType,
		Sender:        req.Sender,
		Target:        req.Target,
		Content:       req.Content,
		Extra:         req.Extra,
		AtUserList:    req.AtUserList,
		IsTransparent: req.IsTransparent,
		ClientUuid:    req.ClientUuid,
	}
	rspL, err := client.GetChatClient().SendMsg(context.Background(), &r)
	if err != nil {
		// TODO
		e := zerrors.FromError(err)
		rsp.Code = e.Code
		rsp.Message = e.Message
		if e.Message == "" {
			rsp.Message = e.Detail
		}
		return
	}

	rsp.Code = rspL.Code
	rsp.Message = rspL.Message
	rsp.Id = rspL.Id
	rsp.SendTime = rspL.SendTime
	rsp.ClientUuid = rspL.ClientUuid
	return
}
