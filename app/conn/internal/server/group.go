package server

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/zmicro-team/zim/app/conn/internal/client"
	"github.com/zmicro-team/zim/app/conn/protocol"
	"github.com/zmicro-team/zim/proto/group"
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
		return
	}

	reqL := group.CreateReq{
		Owner:        req.Owner,
		Members:      req.Members,
		Name:         req.Name,
		GroupId:      req.GroupId,
		Notice:       req.Notice,
		Introduction: req.Introduction,
		Avatar:       req.Avatar,
	}

	rspL, err := client.GetGroupClient().Create(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	rsp.GroupId = rspL.GroupId

	return
}

func (s *Server) handleGetJoinedGroups(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.GetJoinedGroupsReq{}
	rsp := &protocol.GetJoinedGroupsRsp{}

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

	reqL := group.GetJoinedGroupsReq{
		Uin: c.Uin,
	}

	rspL, err := client.GetGroupClient().GetJoinedGroups(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	for _, v := range rspL.Groups {
		g := protocol.GroupInfo{
			Owner:        v.Owner,
			Name:         v.Name,
			GroupId:      v.GroupId,
			Notice:       v.Notice,
			Introduction: v.Introduction,
			Avatar:       v.Avatar,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			Type:         v.Type,
		}
		rsp.Groups = append(rsp.Groups, &g)
	}

	return
}

func (s *Server) handleSyncGroups(c *Client, p *protocol.Packet) (err error) {
	req := &protocol.SyncGroupsReq{}
	rsp := &protocol.SyncGroupsRsp{}

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
			Owner:        v.Owner,
			Name:         v.Name,
			GroupId:      v.GroupId,
			Notice:       v.Notice,
			Introduction: v.Introduction,
			Avatar:       v.Avatar,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			Type:         v.Type,
		}
		rsp.List = append(rsp.List, grp)
	}

	return
}
