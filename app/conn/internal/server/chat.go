package server

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/zchat-team/zim/app/conn/internal/client"
	"github.com/zchat-team/zim/app/conn/protocol"
	"github.com/zchat-team/zim/proto/rpc/chat"
	zerrors "github.com/zmicro-team/zmicro/core/errors"
	"github.com/zmicro-team/zmicro/core/log"
)

func (s *Server) handleMsgAck(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.MsgAckReq{}
	rsp := &protocol.MsgAckRsp{}

	defer func() {
		if err != nil {
			s.responseError(c, p, err)
		} else {
			s.responseMessage(c, p, rsp)
		}
	}()

	if err = proto.Unmarshal(p.Body, req); err != nil {
		return
	}

	reqL := chat.MsgAckReq{
		Uin: c.Uin,
		Id:  req.Id,
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
		if err != nil {
			s.responseError(c, p, err)
		} else {
			s.responseMessage(c, p, rsp)
		}
	}()

	if err = proto.Unmarshal(p.Body, req); err != nil {
		return
	}

	reqL := chat.SyncMsgReq{
		Uin:    c.Uin,
		Offset: req.Offset,
		Limit:  req.Limit,
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
			SendTime:   v.SendTime,
			ClientUuid: v.ClientUuid,
			AtUserList: v.AtUserList,
		}
		rsp.List = append(rsp.List, msg)
	}

	return
}

func (s *Server) responseError(c *Client, p *protocol.Packet, err error) {
	rsp := &protocol.Error{}
	zerr := zerrors.FromError(err)
	rsp.Code = zerr.Code
	rsp.Message = zerr.Message
	if zerr.Message == "" {
		rsp.Message = zerr.Detail
	}
	b, _ := proto.Marshal(rsp)
	p.BodyLen = uint32(len(b))
	p.Body = b
	_ = c.WritePacket(p)
}

func (s *Server) responseMessage(c *Client, p *protocol.Packet, m proto.Message) {
	b, _ := proto.Marshal(m)
	p.BodyLen = uint32(len(b))
	p.Body = b
	_ = c.WritePacket(p)
}

func (s *Server) handleSend(c *Client, p *protocol.Packet) (err error) {
	log.Info("handleSend ...")
	req := &protocol.SendReq{}
	rsp := &protocol.SendRsp{}

	defer func() {
		if err != nil {
			s.responseError(c, p, err)
		} else {
			s.responseMessage(c, p, rsp)
		}
	}()

	if err = proto.Unmarshal(p.Body, req); err != nil {
		log.Error(err)
		return
	}

	r := chat.SendReq{
		ConvType:      req.ConvType,
		MsgType:       req.MsgType,
		Sender:        req.Sender,
		Target:        req.Target,
		Content:       req.Content,
		ClientUuid:    req.ClientUuid,
		AtUserList:    req.AtUserList,
		IsTransparent: req.IsTransparent,
	}
	rspL, err := client.GetChatClient().SendMsg(context.Background(), &r)
	if err != nil {
		log.Error(err)
		return
	}

	rsp.Id = rspL.Id
	rsp.SendTime = rspL.SendTime
	rsp.ClientUuid = r.ClientUuid
	return
}

func (s *Server) handleRecall(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.RecallReq{}
	rsp := &protocol.RecallRsp{}

	defer func() {
		if err != nil {
			s.responseError(c, p, err)
		} else {
			s.responseMessage(c, p, rsp)
		}

	}()

	if err = proto.Unmarshal(p.Body, req); err != nil {
		return
	}

	reqL := chat.RecallReq{
		Uin: c.Uin,
		Id:  req.Id,
	}

	_, err = client.GetChatClient().Recall(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	return
}
