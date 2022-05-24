package service

import (
	"context"
	"github.com/spf13/cast"
	"github.com/zmicro-team/zim/app/chat/internal/model"
	"github.com/zmicro-team/zim/pkg/idgen"
	"github.com/zmicro-team/zim/pkg/runtime"
	"github.com/zmicro-team/zim/proto/group"
	"github.com/zmicro-team/zmicro/core/log"
	"gorm.io/gorm"
)

type Group struct {
}

func GetGroupService() *Group {
	return &Group{}
}

func (g *Group) Create(ctx context.Context, req *group.CreateReq, rsp *group.CreateRsp) (err error) {
	grp := model.Group{
		Id:           idgen.Next(),
		Owner:        req.Owner,
		GroupId:      "",
		Type:         0,
		Name:         req.Name,
		DeletedAt:    0,
		Notice:       req.Notice,
		Introduction: req.Introduction,
		Avatar:       req.Avatar,
	}

	if req.GroupId != "" {
		grp.GroupId = req.GroupId
	} else {
		grp.GroupId = cast.ToString(grp.Id)
	}

	var members []*model.GroupMember
	for _, v := range req.Members {
		member := &model.GroupMember{
			Id:      grp.Id,
			GroupId: grp.GroupId,
			Member:  v,
		}
		members = append(members, member)
	}

	db := runtime.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&g).Error; err != nil {
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

func (g *Group) GetJoinedGroups(ctx context.Context, req *group.GetJoinedGroupsReq, rsp *group.GetJoinedGroupsRsp) (err error) {
	db := runtime.GetDB()
	err = db.Model(&model.GroupMember{}).Where(&model.GroupMember{Member: req.Uin}).
		Select([]string{
			"group.owner",
			"group.group_id",
			"group.type",
			"group.name",
			"UNIX_TIMESTAMP(group.created_at) AS created_at",
			"UNIX_TIMESTAMP(group.updated_at) AS updated_at",
			"group.notice",
			"group.introduction",
			"group.avatar",
		}).
		Joins("INNER JOIN group on group_member.group_id=group.group_id").
		Find(&rsp.Groups).Error

	return
}

func (g *Group) Sync(ctx context.Context, req *group.SyncReq, rsp *group.SyncRsp) (err error) {
	if req.Limit == 0 {
		req.Limit = 20
	} else if req.Limit > 100 {
		req.Limit = 100
	}

	db := runtime.GetDB()
	db = db.Scopes(func(db *gorm.DB) *gorm.DB {
		if req.Offset > 0 {
			db = db.Where("UNIX_TIMESTAMP(group.created_at) > ?", req.Offset)
		}
		return db
	})
	err = db.Model(&model.Group{}).Where(&model.GroupMember{Member: req.Uin}).
		Select([]string{
			"group.owner",
			"group.group_id",
			"group.type",
			"group.name",
			"UNIX_TIMESTAMP(group.created_at) AS created_at",
			"UNIX_TIMESTAMP(group.updated_at) AS updated_at",
			"group.notice",
			"group.introduction",
			"group.avatar",
		}).
		Scopes(func(db *gorm.DB) *gorm.DB {
			return db
		}).
		Joins("INNER JOIN group on group_member.group_id=group.group_id").
		Order("group.updated_at ASC").
		Find(&rsp.List).Error

	return
}
