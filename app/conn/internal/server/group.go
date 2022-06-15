package server

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/zchat-team/zim/app/conn/internal/client"
	"github.com/zchat-team/zim/app/conn/protocol"
	"github.com/zchat-team/zim/proto/rpc/group"
	zerrors "github.com/zmicro-team/zmicro/core/errors"
	"github.com/zmicro-team/zmicro/core/log"
)

func (s *Server) handleCreateGroup(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.CreateGroupReq{}
	rsp := &protocol.CreateGroupRsp{}

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

	reqL := group.CreateReq{
		Owner:   c.Uin,
		Members: req.Members,
		Name:    req.Name,
		GroupId: req.GroupId,
		Notice:  req.Notice,
		Intro:   req.Intro,
		Avatar:  req.Avatar,
	}

	rspL, err := client.GetGroupClient().Create(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	rsp.GroupId = rspL.GroupId

	return
}

func (s *Server) handleGetJoinedGroupList(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.GetJoinedGroupListReq{}
	rsp := &protocol.GetJoinedGroupListRsp{}

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

	reqL := group.GetJoinedGroupListReq{
		Uin: c.Uin,
	}

	rspL, err := client.GetGroupClient().GetJoinedGroupList(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	for _, v := range rspL.List {
		g := protocol.GroupInfo{
			Owner:     v.Owner,
			Name:      v.Name,
			GroupId:   v.GroupId,
			Notice:    v.Notice,
			Intro:     v.Intro,
			Avatar:    v.Avatar,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Type:      v.Type,
		}
		rsp.List = append(rsp.List, &g)
	}

	return
}

func (s *Server) handleSyncGroup(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.SyncGroupReq{}
	rsp := &protocol.SyncGroupRsp{}

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

	reqL := group.SyncReq{
		Uin:      c.Uin,
		DeviceId: c.DeviceId,
		Offset:   req.Offset,
		Limit:    req.Limit,
	}

	rspL, err := client.GetGroupClient().Sync(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	for _, v := range rspL.List {
		grp := &protocol.GroupInfo{
			Owner:     v.Owner,
			Name:      v.Name,
			GroupId:   v.GroupId,
			Notice:    v.Notice,
			Intro:     v.Intro,
			Avatar:    v.Avatar,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Type:      v.Type,
		}
		rsp.List = append(rsp.List, grp)
	}

	return
}

func (s *Server) handleJoinGroup(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.JoinGroupReq{}
	rsp := &protocol.JoinGroupRsp{}

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

	reqL := group.JoinGroupReq{
		Uin:     c.Uin,
		GroupId: req.GroupId,
	}

	_, err = client.GetGroupClient().JoinGroup(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

func (s *Server) handleInviteUserToGroup(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.InviteUserToGroupReq{}
	rsp := &protocol.InviteUserToGroupRsp{}

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

	reqL := group.InviteUserToGroupReq{
		Uin:      c.Uin,
		UserList: req.UserList,
	}

	_, err = client.GetGroupClient().InviteUserToGroup(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

func (s *Server) handleQuitGroup(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.QuitGroupReq{}
	rsp := &protocol.QuitGroupRsp{}

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

	reqL := group.QuitGroupReq{
		Uin:     c.Uin,
		GroupId: req.GroupId,
	}

	_, err = client.GetGroupClient().QuitGroup(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

func (s *Server) handleKickGroupMember(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.KickGroupMemberReq{}
	rsp := &protocol.KickGroupMemberRsp{}

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

	reqL := group.KickGroupMemberReq{
		Uin:      c.Uin,
		GroupId:  req.GroupId,
		UserList: req.UserList,
	}

	_, err = client.GetGroupClient().KickGroupMember(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

func (s *Server) handleDismissGroup(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.DismissGroupReq{}
	rsp := &protocol.DismissGroupRsp{}

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

	reqL := group.DismissGroupReq{
		Uin:     c.Uin,
		GroupId: req.GroupId,
	}

	_, err = client.GetGroupClient().DismissGroup(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

func (s *Server) handleGetGroupMemberList(c *Client, p *protocol.Packet) (err error) {
	return
}

func (s *Server) handleGetGroupMemberInfo(c *Client, p *protocol.Packet) (err error) {
	return
}

func (s *Server) handleSetGroupMemberInfo(c *Client, p *protocol.Packet) (err error) {
	return
}
