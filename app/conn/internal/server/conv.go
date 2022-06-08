package server

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/zchat-team/zim/app/conn/internal/client"
	"github.com/zchat-team/zim/app/conn/protocol"
	"github.com/zchat-team/zim/proto/chat"
	zerrors "github.com/zmicro-team/zmicro/core/errors"
	"github.com/zmicro-team/zmicro/core/log"
)

func (s *Server) handleGetRecentConversation(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.GetRecentConversationReq{}
	rsp := &protocol.GetRecentConversationRsp{}
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

	reqL := chat.GetRecentConversationReq{
		Uin:      c.Uin,
		DeviceId: c.DeviceId,
	}

	rspL, err := client.GetConvClient().GetRecentConversation(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	for _, v := range rspL.List {
		conv := protocol.Conversation{
			Type:   v.Type,
			Target: v.Target,
			IsTop:  v.IsTop,
			IsMute: v.IsMute,
			Remark: v.Remark,
		}
		rsp.List = append(rsp.List, &conv)
	}

	return
}

func (s *Server) handleGetConversationMsg(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.GetConversationMsgReq{}
	rsp := &protocol.GetConversationMsgRsp{}

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

	reqL := chat.GetConversationMsgReq{
		Uin:      c.Uin,
		DeviceId: c.DeviceId,
		ConvId:   req.ConvId,
		Offset:   req.Offset,
		Limit:    req.Limit,
	}

	rspL, err := client.GetConvClient().GetConversationMsg(context.Background(), &reqL)
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

func (s *Server) handleDeleteConversation(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.DeleteConversationReq{}
	rsp := &protocol.DeleteConversationRsp{}

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

	reqL := chat.DeleteConversationReq{
		Uin:      c.Uin,
		DeviceId: c.DeviceId,
		ConvIds:  req.ConvIds,
	}

	_, err = client.GetConvClient().DeleteConversation(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (s *Server) handleGetConversation(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.GetConversationReq{}
	rsp := &protocol.GetConversationRsp{}

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

	reqL := chat.GetConversationReq{
		Uin:      c.Uin,
		DeviceId: c.DeviceId,
		ConvId:   req.ConvId,
	}

	rspL, err := client.GetConvClient().GetConversation(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	rsp.Type = rspL.Type
	rsp.Target = rspL.Target
	rsp.IsTop = rspL.IsTop
	rsp.IsMute = rspL.IsMute
	rsp.Remark = rspL.Remark

	return
}

func (s *Server) handleSetConversationTop(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.SetConversationTopReq{}
	rsp := &protocol.SetConversationTopRsp{}

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

	reqL := chat.SetConversationTopReq{
		Uin:      c.Uin,
		DeviceId: c.DeviceId,
		ConvId:   req.ConvId,
		IsTop:    req.IsTop,
	}

	_, err = client.GetConvClient().SetConversationTop(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (s *Server) handleSetConversationMute(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.SetConversationMuteReq{}
	rsp := &protocol.SetConversationMuteRsp{}

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

	reqL := chat.SetConversationMuteReq{
		Uin:      c.Uin,
		DeviceId: c.DeviceId,
		ConvId:   req.ConvId,
		IsMute:   req.IsMute,
	}

	_, err = client.GetConvClient().SetConversationMute(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (s *Server) handleSyncConversation(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.SyncConversationReq{}
	rsp := &protocol.SyncConversationRsp{}

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

	reqL := chat.SyncConversationReq{
		Uin:      c.Uin,
		DeviceId: c.DeviceId,
		Offset:   req.Offset,
		Limit:    req.Limit,
	}

	rspL, err := client.GetConvClient().SyncConversation(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	for _, v := range rspL.List {
		conv := protocol.Conversation{
			Type:         v.Type,
			Target:       v.Target,
			IsTop:        v.IsTop,
			IsMute:       v.IsMute,
			Remark:       v.Remark,
			PeerLastRead: v.PeerLastRead,
			PeerLastRecv: v.PeerLastRecv,
			UpdatedAt:    v.UpdatedAt,
			UnreadCount:  v.UnreadCount,
			//LastMsg:      v.LastMsg,
		}
		conv.LastMsg = &protocol.Msg{
			Id:         v.LastMsg.Id,
			ConvType:   v.LastMsg.ConvType,
			Type:       v.LastMsg.Type,
			Content:    v.LastMsg.Content,
			Sender:     v.LastMsg.Sender,
			Target:     v.LastMsg.Target,
			SendTime:   v.LastMsg.SendTime,
			ClientUuid: v.LastMsg.ClientUuid,
			AtUserList: v.LastMsg.AtUserList,
		}
		rsp.List = append(rsp.List, &conv)
	}

	return
}

func (s *Server) handleSyncConversationMsg(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.SyncConversationMsgReq{}
	rsp := &protocol.SyncConversationMsgRsp{}

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

	reqL := chat.SyncConversationMsgReq{
		Uin:      c.Uin,
		DeviceId: c.DeviceId,
		ConvId:   req.ConvId,
		Offset:   req.Offset,
		Limit:    req.Limit,
	}

	rspL, err := client.GetConvClient().SyncConversationMsg(context.Background(), &reqL)
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

func (s *Server) handleSetConversationRead(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.SetConversationReadReq{}
	rsp := &protocol.SetConversationReadRsp{}

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

	reqL := chat.SetConversationReadReq{
		Uin:      c.Uin,
		DeviceId: c.DeviceId,
		ConvId:   req.ConvId,
	}

	_, err = client.GetConvClient().SetConversationRead(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	return
}
