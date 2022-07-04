package service

import (
	"context"

	"github.com/spf13/cast"
	"github.com/zmicro-team/zmicro/core/log"
	"gorm.io/gorm"

	"github.com/zchat-team/zim/app/chat/internal/model"
	"github.com/zchat-team/zim/pkg/idgen"
	"github.com/zchat-team/zim/pkg/runtime"
	"github.com/zchat-team/zim/proto/rpc/group"
)

type Group struct {
}

func GetGroupService() *Group {
	return &Group{}
}

func (g *Group) Create(ctx context.Context, req *group.CreateReq, rsp *group.CreateRsp) (err error) {
	grp := model.Group{
		Id:      idgen.Next(),
		Owner:   req.Owner,
		GroupId: "",
		Type:    0,
		Name:    req.Name,
	}

	if req.GroupId != "" {
		grp.GroupId = req.GroupId
	} else {
		grp.GroupId = cast.ToString(grp.Id)
	}

	var members []*model.GroupMember
	members = append(members, &model.GroupMember{
		Id:      idgen.Next(),
		GroupId: grp.GroupId,
		Member:  req.Owner,
	})
	for _, v := range req.Members {
		if v == req.Owner {
			continue
		}
		member := &model.GroupMember{
			Id:      idgen.Next(),
			GroupId: grp.GroupId,
			Member:  v,
		}
		members = append(members, member)
	}

	db := runtime.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&grp).Error; err != nil {
			return err
		}
		if err := tx.Create(&members).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error(err)
		return
	}
	rsp.GroupId = grp.GroupId

	return
}

func (g *Group) GetJoinedGroupList(ctx context.Context, req *group.GetJoinedGroupListReq, rsp *group.GetJoinedGroupListRsp) (err error) {
	db := runtime.GetDB()
	var rows []*model.Group
	err = db.Model(&model.GroupMember{}).Where(&model.GroupMember{Member: req.Uin}).
		Select([]string{
			"group.owner",
			"group.group_id",
			"group.type",
			"group.name",
			"group.created_at",
			"group.updated_at",
		}).
		Joins("INNER JOIN `group` on group_member.group_id=group.group_id").
		Find(&rows).Error

	for _, v := range rows {
		groupInfo := group.GroupInfo{
			Owner:     v.Owner,
			Name:      v.Name,
			GroupId:   v.GroupId,
			CreatedAt: v.CreatedAt.Unix(),
			UpdatedAt: v.UpdatedAt.Unix(),
			Type:      int32(v.Type),
		}
		rsp.List = append(rsp.List, &groupInfo)
	}

	return
}

func (g *Group) Sync(ctx context.Context, req *group.SyncReq, rsp *group.SyncRsp) (err error) {
	if req.Limit == 0 {
		req.Limit = 20
	} else if req.Limit > 100 {
		req.Limit = 100
	}

	db := runtime.GetDB()
	err = db.Model(&model.Group{}).Where(&model.GroupMember{Member: req.Uin}).
		Select([]string{
			"group.owner",
			"group.group_id",
			"group.type",
			"group.name",
			"UNIX_TIMESTAMP(group.created_at) AS created_at",
			"UNIX_TIMESTAMP(group.updated_at) AS updated_at",
		}).
		Scopes(func(db *gorm.DB) *gorm.DB {
			if req.Offset > 0 {
				db = db.Where("UNIX_TIMESTAMP(group.created_at) > ?", req.Offset)
			}
			return db
		}).
		Joins("INNER JOIN group on group_member.group_id=group.group_id").
		Order("group.updated_at ASC").
		Find(&rsp.List).Error

	return
}

func (g *Group) JoinGroup(ctx context.Context, req *group.JoinGroupReq, rsp *group.JoinGroupRsp) (err error) {
	db := runtime.GetDB()
	v := model.GroupMember{}
	err = db.Model(&model.GroupMember{}).Find(&v, &model.GroupMember{GroupId: req.GroupId, Member: req.Uin}).Error
	if v.Id == 0 {
		if err = db.Create(&model.GroupMember{GroupId: req.GroupId, Member: req.Uin}).Error; err != nil {
			log.Error(err)
			return
		}
	}
	return
}

func (g *Group) InviteUserToGroup(ctx context.Context, req *group.InviteUserToGroupReq, rsp *group.InviteUserToGroupRsp) (err error) {
	// TODO: 判断群是否存在
	db := runtime.GetDB()
	var members []*model.GroupMember
	for _, v := range req.UserList {
		member := &model.GroupMember{
			Id:      idgen.Next(),
			GroupId: req.GroupId,
			Member:  v,
		}
		members = append(members, member)
	}

	err = db.Create(&members).Error

	return
}

func (g *Group) QuitGroup(ctx context.Context, req *group.QuitGroupReq, rsp *group.QuitGroupRsp) (err error) {
	return
}

func (g *Group) KickGroupMember(ctx context.Context, req *group.KickGroupMemberReq, rsp *group.KickGroupMemberRsp) (err error) {
	return
}

func (g *Group) DismissGroup(ctx context.Context, req *group.DismissGroupReq, rsp *group.DismissGroupRsp) (err error) {
	return
}

func (g *Group) GetGroupMemberList(ctx context.Context, req *group.GetGroupMemberListReq, rsp *group.GetGroupMemberListRsp) (err error) {
	if req.Limit == 0 {
		req.Limit = 20
	} else if req.Limit > 100 {
		req.Limit = 100
	}

	db := runtime.GetDB()
	err = db.Model(&model.GroupMember{}).Where(&model.GroupMember{GroupId: req.GroupId}).
		Select([]string{
			"group_id",
			"member",
			"UNIX_TIMESTAMP(created_at) AS created_at",
			"UNIX_TIMESTAMP(updated_at) AS updated_at",
		}).
		Scopes(func(db *gorm.DB) *gorm.DB {
			if req.Offset > 0 {
				db = db.Where("UNIX_TIMESTAMP(group_member.updated_at) > ?", req.Offset)
			}
			return db
		}).
		Order("updated_at ASC").
		Find(&rsp.List).Error

	return
}

func (g *Group) GetGroupMemberInfo(ctx context.Context, req *group.GetGroupMemberInfoReq, rsp *group.GetGroupMemberInfoRsp) (err error) {
	//db := runtime.GetDB()
	//v := model.GroupMember{}
	//if err = db.Model(&model.GroupMember{}).
	//	Find(&v, &model.GroupMember{GroupId: req.GroupId, Member: req.Member}).
	//	Error; err != nil {
	//	log.Error(err)
	//	return
	//}
	//if v.Id != 0 {
	//	*rsp = group.GetGroupMemberInfoRsp{
	//		GroupId:   v.GroupId,
	//		Member:    v.Member,
	//		CreatedAt: v.CreatedAt.Unix(),
	//		UpdatedAt: v.UpdatedAt.Unix(),
	//	}
	//}
	return
}

func (g *Group) SetGroupMemberInfo(ctx context.Context, req *group.SetGroupMemberInfoReq, rsp *group.SetGroupMemberInfoRsp) (err error) {
	//db := runtime.GetDB()
	//if err = db.Model(&model.GroupMember{}).
	//	Where(&model.GroupMember{
	//		GroupId: req.GroupId,
	//		Member:  req.Member,
	//	}).
	//	Updates(&model.GroupMember{
	//		Nickname: req.Nickname,
	//	}).Error; err != nil {
	//	log.Error(err)
	//	return
	//}
	return
}
