package contact

import (
	"context"
	"errors"
	"github.com/spf13/cast"
	"github.com/zchat-team/zim/app/demo/internal/constant"
	"github.com/zchat-team/zim/app/demo/internal/model"
	"github.com/zchat-team/zim/app/demo/internal/types"
	"github.com/zchat-team/zim/errno"
	"github.com/zchat-team/zim/pkg/idgen"
	"github.com/zchat-team/zim/pkg/runtime"
	"github.com/zchat-team/zim/pkg/zcontext"
	api "github.com/zchat-team/zim/proto/http/demo/contact"
	"github.com/zmicro-team/zmicro/core/config"
	"github.com/zmicro-team/zmicro/core/log"
	zgin "github.com/zmicro-team/zmicro/core/transport/http"
	"gopkg.in/resty.v1"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Service struct {
	zgin.Implemented
	db *gorm.DB
}

var (
	service *Service
	once    sync.Once
)

func GetService() *Service {
	once.Do(func() {
		service = &Service{}
		service.db = runtime.GetDB()
	})
	return service
}

func (s *Service) Add(ctx context.Context, req *api.AddReq, rsp *api.AddRsp) (err error) {
	uid := zcontext.GetUid(ctx)
	fr := &model.Application{Uid: uid, FriendUid: req.Uid, Status: 1}
	err = s.db.Order("created_at DESC").Where(&fr).
		First(&fr).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}

		if time.Now().Before(time.Unix(fr.ExpiresAt, 0)) {
			// TODO: 是否要更新过期时间
			return nil
		}
	}

	var exists bool
	if err = s.db.Model(&model.Friend{}).Select("count(id)>0").
		Where(&model.Friend{Uid: uid, FriendUid: req.Uid}).
		Find(&exists).Error; err != nil {
		log.Error(err)
		return
	}

	if exists {
		log.Debugf("已经是好友，无需申请 uid=%d friend_uid=%d", uid, req.Uid)
		err = errno.ErrCustom("已经是好友，无需要申请")
		return
	}

	if err = s.db.Create(&model.Application{
		Id:        idgen.Next(),
		Uid:       uid,
		FriendUid: req.Uid,
		Status:    1,
		IsRead:    "0",
		ExpiresAt: time.Now().AddDate(0, 0, 7).Unix(),
	}).Error; err != nil {
		log.Error(err)
		return
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("recover=%v", err)
			}
		}()

		msg := types.Msg{
			ConvType:      1,
			MsgType:       constant.MsgFriendRequest,
			Sender:        "",
			Target:        cast.ToString(req.Uid),
			Content:       "",
			IsTransparent: true,
		}

		client := resty.New()

		restServer := config.GetString("rest")
		sendUrl := restServer + "api/zim/send"
		result, err := client.R().SetBody(&msg).Post(sendUrl)
		if err != nil {
			log.Error(err)
			return
		}

		if !result.IsSuccess() {
			err = errno.ErrCustom("请求失败")
			log.Debug(result.String())
			return
		}

	}()

	return
}

func (s *Service) SetApplicationRead(ctx context.Context, req *api.SetApplicationReadReq, rsp *api.SetApplicationReadRsp) (err error) {
	uid := zcontext.GetUid(ctx)
	if err = s.db.Model(&model.Application{}).
		Where(&model.Application{IsRead: "0", FriendUid: uid}).
		Updates(&model.Application{IsRead: "1"}).Error; err != nil {
		return
	}

	return
}

func (s *Service) GetApplicationList(ctx context.Context, req *api.GetApplicationListReq, rsp *api.GetApplicationListRsp) (err error) {
	uid := zcontext.GetUid(ctx)

	var rows []*api.Application
	if err = s.db.Model(&model.Application{}).Select([]string{
		"application.id", "application.uid", "application.status",
		"UNIX_TIMESTAMP(application.created_at) AS created_at", "UNIX_TIMESTAMP(application.updated_at) AS updated_at",
		"expires_at", "user.nickname", "user.avatar",
	}).Where(&model.Application{FriendUid: uid, IsRead: "0"}).
		Where("expires_at>=?", time.Now().Unix()).
		Joins("LEFT JOIN user ON application.uid=user.id").
		Find(&rows).Error; err != nil {
		log.Error(err)
		return
	}

	rsp.List = rows

	return
}

func (s *Service) List(ctx context.Context, req *api.ListReq, rsp *api.ListRsp) (err error) {
	if req.Limit == 0 {
		req.Limit = 20
	} else if req.Limit > 100 {
		req.Limit = 100
	}

	uid := zcontext.GetUid(ctx)

	var rows []*api.Friend

	if err = s.db.Model(&model.Friend{}).Select([]string{
		"friend.id", "friend.friend_uid",
		"UNIX_TIMESTAMP(friend.created_at) AS created_at", "friend.alias", "user.nickname", "user.avatar",
	}).Where(&model.Friend{Uid: uid}).
		Scopes(func(db *gorm.DB) *gorm.DB {
			if req.Offset > 0 {
				db = db.Where("UNIX_TIMESTAMP(friend.created_at) > ?", req.Offset)
			}
			return db
		}).Joins("LEFT JOIN user ON friend.friend_uid=user.id").
		Limit(int(req.Limit)).
		Find(&rows).Error; err != nil {
		log.Error(err)
		return
	}

	rsp.List = rows

	return
}

func (s *Service) Accept(ctx context.Context, req *api.AcceptReq, rsp *api.AcceptRsp) (err error) {
	uid := zcontext.GetUid(ctx)
	fr := model.Application{Id: req.Id}
	if err = s.db.Take(&fr).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//err = errno.ErrRecordNotFound()
		}
		return
	}

	if fr.FriendUid != uid {
		err = errno.ErrCustom("请求数据异常")
		return
	}

	status := 2
	if err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.db.Updates(&model.Application{Id: req.Id, Status: status}).Error; err != nil {
			return err
		}

		var rows []*model.Friend
		f := model.Friend{
			Id:        idgen.Next(),
			Uid:       uid,
			FriendUid: fr.Uid,
		}
		f2 := model.Friend{
			Id:        idgen.Next(),
			Uid:       fr.Uid,
			FriendUid: uid,
		}
		rows = append(rows, &f, &f2)
		if err := s.db.Create(rows).Error; err != nil {
			log.Error(err)
			return nil
		}
		return nil
	}); err != nil {
		log.Error(err)
		return
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("recover=%v", err)
			}
		}()

		msg := types.Msg{
			ConvType: 1,
			MsgType:  3002,
			Sender:   "",
			Target:   cast.ToString(fr.Uid),
			Content:  "",
		}

		client := resty.New()
		restServer := config.GetString("rest")
		sendUrl := restServer + "api/zim/send"
		result, err := client.R().SetBody(&msg).Post(sendUrl)
		if err != nil {
			log.Error(err)
			return
		}

		if !result.IsSuccess() {
			err = errno.ErrCustom("请求失败")
			log.Debug(result.String())
			return
		}
	}()

	return
}

func (s *Service) Refuse(ctx context.Context, req *api.RefuseReq, rsp *api.RefuseRsp) (err error) {
	uid := zcontext.GetUid(ctx)
	fr := model.Application{Id: req.Id}
	if err = s.db.Take(&fr).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//err = errno.ErrRecordNotFound()
		}
		return
	}

	if fr.FriendUid != uid {
		err = errno.ErrCustom("请求数据异常")
		return
	}

	status := 3
	if err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.db.Updates(&model.Application{Id: req.Id, Status: status}).Error; err != nil {
			return err
		}

		var rows []*model.Friend
		f := model.Friend{
			Id:        idgen.Next(),
			Uid:       uid,
			FriendUid: fr.Uid,
		}
		f2 := model.Friend{
			Id:        idgen.Next(),
			Uid:       fr.Uid,
			FriendUid: uid,
		}
		rows = append(rows, &f, &f2)
		if err := s.db.Create(rows).Error; err != nil {
			log.Error(err)
			return nil
		}
		return nil
	}); err != nil {
		log.Error(err)
		return
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("recover=%v", err)
			}
		}()

		msg := types.Msg{
			ConvType:      1,
			MsgType:       3002,
			Sender:        "",
			Target:        cast.ToString(uid),
			Content:       "",
			IsTransparent: true,
		}

		client := resty.New()
		restServer := config.GetString("rest")
		sendUrl := restServer + "api/zim/send"
		result, err := client.R().SetBody(&msg).Post(sendUrl)
		if err != nil {
			log.Error(err)
			return
		}

		if !result.IsSuccess() {
			err = errno.ErrCustom("请求失败")
			log.Debug(result.String())
			return
		}
	}()

	return
}
