package server

import (
	"context"
	"github.com/zmicro-team/zim/app/conn/internal/client"
	"github.com/zmicro-team/zim/proto/chat"
	"github.com/zmicro-team/zim/proto/conn"

	"github.com/golang/protobuf/proto"
	"github.com/zmicro-team/zim/app/conn/protocol"
	"github.com/zmicro-team/zmicro/core/log"
)

func (s *Server) handleGetRecentConversation(c *Client, p *protocol.Packet) (err error) {
	req := &conn.GetRecentConversationReq{}
	rsp := &conn.GetRecentConversationRsp{}

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
		conv := conn.Conversation{
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
	req := &conn.GetConversationMsgReq{}
	rsp := &conn.GetConversationMsgRsp{}

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
		msg := &conn.Msg{
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
		}
		rsp.List = append(rsp.List, msg)
	}

	return
}

func (s *Server) handleDeleteConversation(c *Client, p *protocol.Packet) (err error) {
	req := &conn.DeleteConversationReq{}
	rsp := &conn.DeleteConversationRsp{}

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
	req := &conn.GetConversationReq{}
	rsp := &conn.GetConversationRsp{}

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
	req := &conn.SetConversationTopReq{}
	rsp := &conn.SetConversationTopRsp{}

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
	req := &conn.SetConversationMuteReq{}
	rsp := &conn.SetConversationMuteRsp{}

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
	req := &conn.SyncConversationReq{}
	rsp := &conn.SyncConversationRsp{}

	defer func() {
		b, err := proto.Marshal(rsp)
		if err != nil {
			return
		}

		//x := protocol.SyncConversationRsp{}
		//err = proto.Unmarshal(b, &x)
		//if err != nil {
		//	logger.Error(err)
		//} else {
		//	logger.Error("kadfadfadfafdadf")
		//}
		//logger.Errorf("bbbbbbbbbb=%s len=%d", string(b), len(b))
		p.BodyLen = uint32(len(b))
		p.Body = b
		c.WritePacket(p)
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
		conv := conn.Conversation{
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
		conv.LastMsg = &conn.Msg{
			Id:            v.LastMsg.Id,
			ConvType:      v.LastMsg.ConvType,
			Type:          v.LastMsg.Type,
			Content:       v.LastMsg.Content,
			Sender:        v.LastMsg.Sender,
			Target:        v.LastMsg.Target,
			Extra:         v.LastMsg.Extra,
			SendTime:      v.LastMsg.SendTime,
			AtUserList:    v.LastMsg.AtUserList,
			ReadTime:      v.LastMsg.ReadTime,
			ClientUuid:    v.LastMsg.ClientUuid,
			IsTransparent: v.LastMsg.IsTransparent,
		}
		rsp.List = append(rsp.List, &conv)
	}

	return
}

func (s *Server) handleSyncConversationMsg(c *Client, p *protocol.Packet) (err error) {
	req := &conn.SyncConversationMsgReq{}
	rsp := &conn.SyncConversationMsgRsp{}

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
		msg := &conn.Msg{
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
		}
		rsp.List = append(rsp.List, msg)
	}

	return
}

func (s *Server) handleSetConversationRead(c *Client, p *protocol.Packet) (err error) {
	req := &conn.SetConversationReadReq{}
	rsp := &conn.SetConversationReadRsp{}

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
