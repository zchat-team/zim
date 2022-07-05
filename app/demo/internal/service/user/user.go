package user

import (
	"context"
	"sync"

	"gorm.io/gorm"

	"github.com/zchat-team/zim/app/demo/internal/model"
	"github.com/zchat-team/zim/pkg/runtime"
	api "github.com/zchat-team/zim/proto/http/demo/user"
	"github.com/zmicro-team/zmicro/core/log"
	zgin "github.com/zmicro-team/zmicro/core/transport/http"
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

func (s *Service) Search(ctx context.Context, req *api.SearchReq, rsp *api.SearchRsp) (err error) {
	var users []*api.UserInfo
	if err = s.db.Model(&model.User{}).
		Select([]string{"id", "zid", "nickname", "avatar"}).
		Where("zid=? OR mobile=?", req.Id, req.Id).
		Find(&users).Error; err != nil {
		return
	}
	rsp.List = users
	return
}

func (s *Service) Get(ctx context.Context, req *api.GetReq, rsp *api.GetRsp) (err error) {
	v := model.User{Id: req.Uid}
	if err = s.db.Take(&v).Error; err != nil {
		log.Error(err)
		return
	}

	*rsp = api.GetRsp{
		Id:       v.Id,
		Zid:      v.Zid,
		Nickname: v.Nickname,
		Avatar:   v.Avatar,
	}

	return
}

func (s *Service) MGet(ctx context.Context, req *api.MGetReq, rsp *api.MGetRsp) (err error) {
	if err = s.db.Model(&model.User{}).
		Select([]string{"id", "zid", "nickname", "avatar"}).
		Find(&rsp.List, req.Uids).Error; err != nil {
		return err
	}

	return
}
