package server

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/zchat-team/zim/app/conn/internal/client"
	"github.com/zchat-team/zim/app/conn/protocol"
	"github.com/zchat-team/zim/pkg/idgen"
	"github.com/zchat-team/zim/pkg/runtime"
	"github.com/zchat-team/zim/proto/rpc/chat"
	"github.com/zchat-team/zim/proto/rpc/common"
	zerrors "github.com/zmicro-team/zmicro/core/errors"
	"github.com/zmicro-team/zmicro/core/log"
	"time"
)

func (s *Server) handleMsgAck(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.MsgAckReq{}
	rsp := &protocol.MsgAckRsp{}

	defer func() {
		var b []byte
		var errr error

		if err != nil {
			rspErr := &protocol.Error{}
			ze := zerrors.FromError(err)
			rspErr.Code = ze.Code
			rspErr.Message = ze.Message
			if ze.Message == "" {
				rspErr.Message = ze.Detail
			}
			b, errr = proto.Marshal(rspErr)
		} else {
			b, errr = proto.Marshal(rsp)
		}

		if errr != nil {
			log.Error(err)
		} else {
			p.BodyLen = uint32(len(b))
			p.Body = b
			if err := c.WritePacket(p); err != nil {
				log.Error(err)
			}
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
		var b []byte
		var errr error

		if err != nil {
			rspErr := &protocol.Error{}
			ze := zerrors.FromError(err)
			rspErr.Code = ze.Code
			rspErr.Message = ze.Message
			if ze.Message == "" {
				rspErr.Message = ze.Detail
			}
			b, errr = proto.Marshal(rspErr)
		} else {
			b, errr = proto.Marshal(rsp)
		}

		if errr != nil {
			log.Error(err)
		} else {
			p.BodyLen = uint32(len(b))
			p.Body = b
			if err := c.WritePacket(p); err != nil {
				log.Error(err)
			}
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

func (s *Server) handleSend(c *Client, p *protocol.Packet) (err error) {
	log.Info("handleSend ...")
	req := &protocol.SendReq{}

	rsp := &protocol.SendRsp{}

	defer func() {
		var b []byte
		var errr error

		if err != nil {
			rspErr := &protocol.Error{}
			ze := zerrors.FromError(err)
			rspErr.Code = ze.Code
			rspErr.Message = ze.Message
			if ze.Message == "" {
				rspErr.Message = ze.Detail
			}
			b, errr = proto.Marshal(rspErr)
		} else {
			b, errr = proto.Marshal(rsp)
		}

		if errr != nil {
			log.Error(err)
		} else {
			p.BodyLen = uint32(len(b))
			p.Body = b
			if err := c.WritePacket(p); err != nil {
				log.Error(err)
			}
		}
	}()

	if err = proto.Unmarshal(p.Body, req); err != nil {
		log.Error(err)
		return
	}

	now := time.Now().UnixNano() / 1e6
	m := common.Msg{
		Id:            idgen.Next(),
		ConvType:      req.ConvType,
		Type:          req.MsgType,
		Content:       req.Content,
		Sender:        req.Sender,
		Target:        req.Target,
		SendTime:      now,
		ClientUuid:    req.ClientUuid,
		AtUserList:    req.AtUserList,
		Owner:         "",
		IsTransparent: req.IsTransparent,
	}

	b, err := proto.Marshal(&m)
	if err != nil {
		return
	}
	nm := &nats.Msg{
		Subject: "MSGS.new",
		Reply:   "",
		Data:    b,
		Sub:     nil,
	}
	js := runtime.GetJS()
	if _, err = js.PublishMsg(nm); err != nil {
		return
	}

	//r := chat.SendReq{
	//	ConvType:      req.ConvType,
	//	MsgType:       req.MsgType,
	//	Sender:        req.Sender,
	//	Target:        req.Target,
	//	Content:       req.Content,
	//	ClientUuid:    req.ClientUuid,
	//	AtUserList:    req.AtUserList,
	//	IsTransparent: req.IsTransparent,
	//}
	//rspL, err := client.GetChatClient().SendMsg(context.Background(), &r)
	//if err != nil {
	//	log.Error(err)
	//	return
	//}

	rsp.Id = m.Id
	rsp.SendTime = m.SendTime
	rsp.ClientUuid = m.ClientUuid
	return
}
