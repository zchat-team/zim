package group

import (
	"encoding/json"
	"github.com/zchat-team/zim/app/demo/internal/constant"
	"github.com/zchat-team/zim/app/demo/internal/model"
	"github.com/zchat-team/zim/app/demo/internal/types"
	"github.com/zchat-team/zim/errno"
	"github.com/zchat-team/zim/pkg/idgen"
	"github.com/zchat-team/zim/pkg/runtime"
	"github.com/zchat-team/zim/pkg/zcontext"
	api "github.com/zchat-team/zim/proto/http/demo/group"
	"github.com/zmicro-team/zmicro/core/log"
	zgin "github.com/zmicro-team/zmicro/core/transport/http"

	"context"
	"github.com/spf13/cast"
	"gopkg.in/resty.v1"
	"gorm.io/gorm"
	"sync"
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

func (s *Service) Create(ctx context.Context, req *api.CreateReq, rsp *api.CreateRsp) (err error) {
	uid := zcontext.GetUid(ctx)
	g := model.Group{
		Id:     idgen.Next(),
		Owner:  uid,
		Type:   0,
		Name:   req.Name,
		Notice: req.Notice,
		Intro:  req.Intro,
		Avatar: req.Avatar,
	}

	var members []*model.GroupMember
	members = append(members, &model.GroupMember{
		Id:      idgen.Next(),
		GroupId: g.Id,
		Member:  uid,
	})
	for _, v := range req.Members {
		if v == uid {
			continue
		}
		m := model.GroupMember{
			Id:      idgen.Next(),
			GroupId: g.Id,
			Member:  v,
		}
		members = append(members, &m)
	}

	if err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.db.Create(&g).Error; err != nil {
			return err
		}

		if len(members) > 0 {
			if err := s.db.Create(&members).Error; err != nil {
				return err
			}
		}

		createUrl := "http://localhost:5080/api/zim/groups"
		client := resty.New()
		createReq := types.CreateReq{
			Owner:   cast.ToString(uid),
			Members: cast.ToStringSlice(req.Members),
			Name:    req.Name,
			GroupId: cast.ToString(g.Id),
			Notice:  "",
			Intro:   "",
			Avatar:  req.Avatar,
		}
		result, err := client.R().SetBody(&createReq).Post(createUrl)
		if err != nil {
			log.Error(err)
			return err
		}
		if !result.IsSuccess() {
			err = errno.ErrCustom("请求失败")
			log.Debug(result.String())
			return err
		}
		createRsp := types.CreateRsp{}
		json.Unmarshal(result.Body(), &createRsp)
		rsp.GroupId = g.Id

		// TODO: 发送通知
		go func() {
			var ids []int64
			for _, v := range members {
				ids = append(ids, cast.ToInt64(v.Member))
			}
			var users []*types.Target
			s.db.Model(&model.User{}).Select([]string{"id", "nickname AS name"}).Find(&users, ids)
			msg := types.Msg{
				ConvType: 2,
				MsgType:  constant.GroupNotify,
				Sender:   "",
				Target:   cast.ToString(g.Id),
				Content:  "",
			}

			data := make(map[string]interface{})
			data["operator_name"] = users[0].Name
			data["target_list"] = users[1:]
			b, _ := json.Marshal(data)
			notify := types.GroupNotify{
				Operator:  cast.ToString(uid),
				Operation: "add",
				Data:      string(b),
			}

			b, _ = json.Marshal(notify)
			msg.Content = string(b)

			sendUrl := "http://localhost:5080/api/zim/send"
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
		return nil
	}); err != nil {
		log.Error(err)
		return
	}

	return
}

func (s *Service) Add(ctx context.Context, req *api.AddReq, rsp *api.AddRsp) (err error) {
	uid := zcontext.GetUidStr(ctx)

	var members []*model.GroupMember

	for _, v := range req.Members {
		m := model.GroupMember{
			Id:      idgen.Next(),
			GroupId: req.GroupId,
			Member:  v,
		}
		members = append(members, &m)
	}

	if len(members) > 0 {
		if err := s.db.Create(&members).Error; err != nil {
			return err
		}
	}

	groupId := cast.ToString(req.GroupId)
	addUrl := "http://localhost:5080/api/zim/groups/" + groupId + "/members"
	client := resty.New()
	addReq := types.AddReq{
		GroupId: groupId,
		Members: cast.ToStringSlice(req.Members),
	}
	result, err := client.R().SetBody(&addReq).Post(addUrl)
	if err != nil {
		log.Error(err)
		return err
	}
	if !result.IsSuccess() {
		err = errno.ErrCustom("请求失败")
		log.Debug(result.String())
		return err
	}

	// TODO: 发送通知
	go func() {
		var ids []int64
		ids = append(ids, cast.ToInt64(uid))
		for _, v := range members {
			ids = append(ids, cast.ToInt64(v.Member))
		}
		var users []*types.Target
		s.db.Model(&model.User{}).Select([]string{"id", "nickname AS name"}).Find(&users, ids)
		msg := types.Msg{
			ConvType: 2,
			MsgType:  constant.GroupNotify,
			Sender:   "",
			Target:   groupId,
			Content:  "",
		}

		data := make(map[string]interface{})
		data["operator_name"] = users[0].Name
		data["target_list"] = users[1:]
		b, _ := json.Marshal(data)
		notify := types.GroupNotify{
			Operator:  uid,
			Operation: "add",
			Data:      string(b),
		}

		b, _ = json.Marshal(notify)
		msg.Content = string(b)

		sendUrl := "http://localhost:5080/api/zim/send"
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

func (s *Service) Join(ctx context.Context, req *api.JoinReq, rsp *api.JoinRsp) (err error) {

	uid := zcontext.GetUidStr(ctx)
	joinReq := types.JoinReq{
		Uin:     uid,
		GroupId: cast.ToString(req.GroupId),
	}
	joinUrl := "http://localhost:5080/api/group/join"
	client := resty.New()
	result, err := client.R().SetBody(&joinReq).Post(joinUrl)
	if err != nil {
		log.Error(err)
		return err
	}
	if !result.IsSuccess() {
		err = errno.ErrCustom("请求失败")
		log.Debug(result.String())
		return err
	}
	return
}

func (s *Service) Quit(ctx context.Context, req *api.QuitReq, rsp *api.QuitRsp) (err error) {
	return
}

func (s *Service) Kick(ctx context.Context, req *api.KickReq, rsp *api.KickRsp) (err error) {
	return
}

func (s *Service) Dismiss(ctx context.Context, req *api.DismissReq, rsp *api.DismissRsp) (err error) {
	return
}

func (s *Service) Transfer(ctx context.Context, req *api.TransferReq, rsp *api.TransferRsp) (err error) {
	return
}

func (s *Service) AddManager(ctx context.Context, req *api.AddManagerReq, rsp *api.AddManagerRsp) (err error) {
	return
}

func (s *Service) RemoveManager(ctx context.Context, req *api.RemoveManagerReq, rsp *api.RemoveManagerRsp) (err error) {
	return
}

func (s *Service) Rename(ctx context.Context, req *api.RenameReq, rsp *api.RenameRsp) (err error) {
	return
}

func (s *Service) SetAvatar(ctx context.Context, req *api.SetAvatarReq, rsp *api.SetAvatarRsp) (err error) {
	return
}

func (s *Service) SetDisplayName(ctx context.Context, req *api.SetDisplayNameReq, rsp *api.SetDisplayNameRsp) (err error) {
	return
}

func (s *Service) MemberList(ctx context.Context, req *api.MemberListReq, rsp *api.MemberListRsp) (err error) {
	// TODO: 获取用户头像
	//var members *model.GroupMember
	//if err = s.db.Where(&model.GroupMember{GroupId: req.GroupId}).
	//	Where("id>?", req.Offset).
	//	Order("ORDER BY id").
	//	Limit(int(req.Limit)).
	//	Find(&members).Error; err != nil {
	//	return
	//}
	return
}

func (s *Service) Info(ctx context.Context, req *api.InfoReq, rsp *api.InfoRsp) (err error) {
	g := model.Group{}
	if err = s.db.Model(&model.Group{}).Find(&g, model.Group{Id: req.GroupId}).Limit(1).Error; err != nil {
		return
	}

	rsp.GroupId = req.GroupId
	rsp.Name = g.Name
	rsp.Notice = g.Notice
	rsp.Avatar = g.Avatar

	return
}
