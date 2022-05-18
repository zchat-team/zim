package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/zmicro-team/zim/pkg/constant"
	"github.com/zmicro-team/zim/pkg/runtime"
	"github.com/zmicro-team/zim/pkg/util"
	"github.com/zmicro-team/zim/proto/common"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cast"
	"github.com/zmicro-team/zmicro/core/log"

	"github.com/zmicro-team/zim/app/chat/internal/model"
	"github.com/zmicro-team/zim/app/chat/internal/typ"
	"github.com/zmicro-team/zim/proto/chat"
)

type Conv struct {
}

func GetConvService() *Conv {
	return &Conv{}
}

// 获取最近的100个会话
func (l *Conv) GetRecentConversation(ctx context.Context, req *chat.GetRecentConversationReq, rsp *chat.GetRecentConversationRsp) (err error) {
	return
}

func (l *Conv) GetConversationMsg(ctx context.Context, req *chat.GetConversationMsgReq, rsp *chat.GetConversationMsgRsp) (err error) {
	// 保证消息库 与 同步库 数据一致
	//if err = l.removeDirty(ctx, req); err != nil {
	//	return
	//}

	if req.Direction == 0 {
		// TODO ======================
		//err = errno.ErrCustom("direction不能为空")
		return
	}

	arr := strings.Split(req.ConvId, "#")
	if len(arr) != 2 {
		err = errors.New("参数错误")
		return
	}
	if req.Limit <= 0 {
		req.Limit = 20
	}
	if req.Limit > 100 {
		req.Limit = 100
	}

	var (
		condition string
		args      []interface{}
	)

	if arr[0] == "C2C" {
		condition = "(sender=? AND target=?) OR (sender=? AND target=?)"
		args = append(args, req.Uin, arr[1], arr[1], req.Uin)
	} else {
		condition = "(sender=? AND target=?)"
		args = append(args, req.Uin, arr[1])
	}

	var order string
	if req.Direction == 1 {
		order = "id DESC"
		if req.Offset > 0 {
			condition += " AND id < ?"
			args = append(args, req.Offset)
		}
	} else if req.Direction == 2 {
		order = "id ASC"
		if req.Offset > 0 {
			condition += " AND id > ?"
			args = append(args, req.Offset)
		}
	}

	var rows []*model.Msg
	db := runtime.GetDB()
	if err = db.Model(&model.Msg{}).
		Order(order).
		Where(condition, args...).
		Limit(int(req.Limit)).
		Find(&rows).Error; err != nil {
		log.Error(err.Error())
		return
	}

	*rsp = chat.GetConversationMsgRsp{}
	for _, v := range rows {
		msg := common.Msg{
			Id:            v.Id,
			ConvType:      int32(v.ConvType),
			Type:          int32(v.Type),
			Content:       v.Content,
			Sender:        v.Sender,
			Target:        v.Target,
			Extra:         v.Extra,
			SendTime:      v.SendTime,
			ReadTime:      v.ReadTime,
			ClientUuid:    v.ClientUuid,
			IsTransparent: false,
		}
		if v.AtUserList != "" {
			json.Unmarshal([]byte(v.AtUserList), &msg.AtUserList)
		}
		rsp.List = append(rsp.List, &msg)
	}

	return
}

func (l *Conv) DeleteConversation(ctx context.Context, req *chat.DeleteConversationReq, rsp *chat.DeleteConversationRsp) (err error) {
	rc := runtime.GetRedisClient()
	var members []interface{}
	for _, v := range req.ConvIds {
		arr := strings.Split(v, "#")
		if len(arr) != 2 {
			err = errors.New("参数错误")
			return
		}
		members = append(members, arr[1])

		rc.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			key := util.KeyConvMsgSync(req.Uin, arr[1])
			pipe.Del(ctx, key)
			key = util.KeyConv(req.Uin, arr[1])
			pipe.Del(ctx, key)
			key = util.KeyConvSync(req.Uin)
			pipe.ZRem(ctx, fmt.Sprintf(key, req.Uin), arr[1])

			return nil
		})
	}

	return
}

func (l *Conv) GetConversation(ctx context.Context, req *chat.GetConversationReq, rsp *common.Conversation) (err error) {
	arr := strings.Split(req.ConvId, "#")
	if len(arr) != 2 {
		err = errors.New("参数错误")
		return
	}

	key := util.KeyConv(req.Uin, arr[1])
	reply, err := l.getConversation(ctx, key)
	*rsp = *reply

	return
}

func (l *Conv) getConversation(ctx context.Context, key string) (conv *common.Conversation, err error) {
	rc := runtime.GetRedisClient()

	v, err := rc.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			return
		}
	}

	conv = &common.Conversation{}
	json.Unmarshal([]byte(v), conv)

	return
}

func (l *Conv) SetConversationTop(ctx context.Context, req *chat.SetConversationTopReq, rsp *chat.SetConversationTopRsp) (err error) {
	arr := strings.Split(req.ConvId, "#")
	if len(arr) != 2 {
		err = errors.New("参数错误")
		return
	}

	key := util.KeyConv(req.Uin, arr[1])
	conv, err := l.getConversation(ctx, key)
	if err != nil {
		log.Error(err.Error())
		return
	}
	conv.IsTop = req.IsTop
	conv.UpdatedAt = time.Now().UnixNano() / 1e6
	b, err := json.Marshal(conv)
	if err != nil {
		return
	}

	rc := runtime.GetRedisClient()
	if err = rc.SetEX(ctx, key, string(b),
		time.Duration(constant.ConvKeepDays*24)*time.Hour).Err(); err != nil {
		return
	}
	//*rsp = api.SetConversationTopRsp{}
	// TODO: 同步给其他端

	return
}

func (l *Conv) SetConversationMute(ctx context.Context, req *chat.SetConversationMuteReq, rsp *chat.SetConversationMuteRsp) (err error) {
	arr := strings.Split(req.ConvId, "#")
	if len(arr) != 2 {
		err = errors.New("参数错误")
		return
	}

	key := util.KeyConv(req.Uin, arr[1])
	conv, err := l.getConversation(ctx, key)
	if err != nil {
		log.Error(err.Error())
		return
	}
	conv.IsMute = req.IsMute
	b, err := json.Marshal(conv)
	if err != nil {
		return
	}

	rc := runtime.GetRedisClient()
	if err = rc.SetEX(ctx, key, string(b),
		time.Duration(constant.ConvKeepDays*24)*time.Hour).Err(); err != nil {
		return
	}

	rsp = &chat.SetConversationMuteRsp{}
	// TODO: 同步给其他端

	return
}

func (l *Conv) SetConversationRead(ctx context.Context, req *chat.SetConversationReadReq, rsp *chat.SetConversationReadRsp) (err error) {
	arr := strings.Split(req.ConvId, "#")
	if len(arr) != 2 {
		err = errors.New("参数错误")
		return
	}
	convType := constant.ConvTypeC2C
	if arr[0] == "C2G" {
		convType = constant.ConvTypeGroup
	}
	now := time.Now().Unix() / 1e6
	if convType == constant.ConvTypeC2C {
		key := util.KeyConv(req.Uin, arr[1])
		conv, err := l.getConversation(ctx, key)
		if err != nil {
			log.Error(err.Error())
			return err
		}

		conv.PeerLastRead = now
		conv.UpdatedAt = now
		b, err := json.Marshal(conv)
		if err != nil {
			return err
		}

		rc := runtime.GetRedisClient()
		if err = rc.SetEX(ctx, key, string(b),
			time.Duration(constant.ConvKeepDays*24)*time.Hour).Err(); err != nil {
			return err
		}
	}

	rsp = &chat.SetConversationReadRsp{}

	// 推送已送达回执
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("err=%v", err)
			}
		}()

		mdr := typ.MsgReadReceipt{
			Uin:         req.Uin,
			ReceiptTime: now,
		}
		b, e := json.Marshal(mdr)
		if e != nil {
			return
		}

		p := common.Msg{
			Id:            0,
			ConvType:      0,
			Type:          constant.MsgReadReceipt,
			Content:       string(b),
			Sender:        "",
			Target:        "",
			Extra:         "",
			SendTime:      now,
			IsTransparent: true,
		}

		b, e = proto.Marshal(&p)
		if e != nil {
			return
		}

		m := &nats.Msg{
			Subject: "MSGS.received",
			Reply:   "",
			Data:    b,
			Sub:     nil,
		}
		js := runtime.GetJS()
		js.PublishMsg(m)
	}()

	return
}

func (l *Conv) SyncConversation(ctx context.Context, req *chat.SyncConversationReq, rsp *chat.SyncConversationRsp) (err error) {
	log.Infof("SyncConversation=%v", req)
	if err = l.removeDirty(ctx, req); err != nil {
		log.Error(err.Error())
		return
	}

	if req.Limit <= 0 {
		req.Limit = 20
	}
	if req.Limit > 100 {
		req.Limit = 100
	}

	zr := redis.ZRangeBy{
		Min:    "(" + cast.ToString(req.Offset),
		Max:    "+inf",
		Offset: 0,
		Count:  req.Limit,
	}

	rc := runtime.GetRedisClient()

	key := util.KeyConvSync(req.Uin)
	cmd := rc.ZRangeByScore(ctx, key, &zr)
	val, err := cmd.Result()
	if err != nil {
		log.Error(err.Error())
		return
	}

	var keys []string
	for _, v := range val {
		key = util.KeyConv(req.Uin, v)
		keys = append(keys, key)
	}

	rsp = &chat.SyncConversationRsp{}
	if len(keys) == 0 {
		return
	}
	rr, err := rc.MGet(ctx, keys...).Result()
	for _, v := range rr {
		if v == nil {
			continue
		}

		conv := common.Conversation{}
		if err := json.Unmarshal([]byte(v.(string)), &conv); err != nil {
			continue
		}

		log.Infof("key=%s", v.(string))
		rsp.List = append(rsp.List, &conv)
	}

	log.Infof("rsp=%v", rsp)
	return
}

func (l *Conv) SyncConversationMsg(ctx context.Context, req *chat.SyncConversationMsgReq, rsp *chat.SyncConversationMsgRsp) (err error) {

	arr := strings.Split(req.ConvId, "#")
	if len(arr) != 2 {
		err = errors.New("参数错误")
		return
	}

	if req.Limit <= 0 {
		req.Limit = 20
	}
	if req.Limit > 100 {
		req.Limit = 100
	}

	zr := redis.ZRangeBy{
		Min:    "(" + cast.ToString(req.Offset),
		Max:    "+inf",
		Offset: 0,
		Count:  req.Limit,
	}

	rc := runtime.GetRedisClient()

	key := util.KeyConvMsgSync(req.Uin, arr[1])
	cmd := rc.ZRangeByScore(ctx, key, &zr)
	val, err := cmd.Result()
	if err != nil {
		log.Error(err.Error())
		return
	}

	var first, second string
	if arr[1] == "C2C" {
		if req.Uin < arr[1] {
			first = req.Uin
			second = arr[1]
		} else {
			first = arr[1]
			second = req.Uin
		}
	} else {
		first = req.Uin
		second = arr[1]
	}

	var keys []string
	for _, v := range val {
		key = util.KeyConvMsg(first, second, cast.ToInt64(v))
		keys = append(keys, key)
	}

	rsp = &chat.SyncConversationMsgRsp{}
	if len(keys) == 0 {
		return
	}
	rr, err := rc.MGet(ctx, keys...).Result()
	for _, v := range rr {
		if v == nil {
			continue
		}
		msg := common.Msg{}
		if err := json.Unmarshal([]byte(v.(string)), &msg); err != nil {
			continue
		}
		rsp.List = append(rsp.List, &msg)
	}

	return
}

func (l *Conv) removeDirty(ctx context.Context, req *chat.SyncConversationReq) (err error) {
	rc := runtime.GetRedisClient()
	// 删除过期数据
	t := time.Now().AddDate(0, 0, -constant.ConvKeepDays)
	max := t.UnixNano() / 1e6
	key := util.KeyConvSync(req.Uin)
	_, err = rc.ZRemRangeByScore(ctx, key, "-inf", cast.ToString(max)).Result()
	if err != nil {
		return
	}

	return
}
