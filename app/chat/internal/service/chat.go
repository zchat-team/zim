package service

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cast"
	"github.com/zchat-team/zim/app/chat/internal/model"
	"github.com/zchat-team/zim/pkg/constant"
	"github.com/zchat-team/zim/pkg/idgen"
	"github.com/zchat-team/zim/pkg/runtime"
	"github.com/zchat-team/zim/pkg/util"
	"github.com/zchat-team/zim/proto/rpc/chat"
	"github.com/zchat-team/zim/proto/rpc/common"
	"github.com/zmicro-team/zmicro/core/log"
	"gorm.io/gorm"
)

type Chat struct {
}

func GetChatService() *Chat {
	return &Chat{}
}

func (l *Chat) SendMsg(ctx context.Context, req *chat.SendReq, rsp *chat.SendRsp) (err error) {
	log.Infof("Chat SendMsg ConvType=%d Type=%d Content=%s", req.ConvType, req.MsgType, req.Content)
	now := time.Now().UnixMilli()
	m := common.Msg{
		Id:            idgen.Next(),
		ConvType:      req.ConvType,
		Type:          req.MsgType,
		Content:       req.Content,
		Sender:        req.Sender,
		Target:        req.Target,
		SendTime:      now,
		ClientUuid:    req.ClientUuid,
		AtUserList:    req.AtUserList,
		Owner:         "",
		IsTransparent: req.IsTransparent,
	}

	b, err := proto.Marshal(&m)
	if err != nil {
		return
	}
	nm := &nats.Msg{
		Subject: "MSGS.new",
		Reply:   "",
		Data:    b,
		Sub:     nil,
	}
	js := runtime.GetJS()
	if _, err = js.PublishMsg(nm); err != nil {
		return
	}

	rsp.Id = m.Id
	rsp.SendTime = m.SendTime
	rsp.ClientUuid = m.ClientUuid
	//if req.ConvType == constant.ConvTypeC2C {
	//	err = l.sendC2C(ctx, req, rsp)
	//} else if req.ConvType == constant.ConvTypeGroup {
	//	err = l.sendC2G(ctx, req, rsp)
	//}
	return
}

// SyncMsg 同步离线消息，从redis缓存中读取，只同步最近30天的消息
func (l *Chat) SyncMsg(ctx context.Context, req *chat.SyncMsgReq, rsp *chat.SyncMsgRsp) (err error) {
	// 保证消息库 与 同步库 数据一致
	if err = l.removeDirty(ctx, req); err != nil {
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

	key := util.KeyMsgSync(req.Uin)
	cmd := rc.ZRangeByScore(ctx, key, &zr)
	val, err := cmd.Result()
	if err != nil {
		return
	}

	var keys []string
	for _, v := range val {
		key = util.KeyMsg(req.Uin, cast.ToInt64(v))
		keys = append(keys, key)
	}

	//*rsp = api.SyncMsgRsp{}
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

func (l *Chat) removeDirty(ctx context.Context, req *chat.SyncMsgReq) (err error) {
	rc := runtime.GetRedisClient()
	// 删除过期id
	t := time.Now().AddDate(0, 0, -constant.MsgKeepDays)
	max := t.UnixNano() / 1e6
	key := util.KeyMsgSync(req.Uin)
	_, err = rc.ZRemRangeByScore(ctx, key, "-inf", cast.ToString(max)).Result()
	if err != nil {
		return
	}

	for {
		zr := redis.ZRangeBy{
			Min:    "(" + cast.ToString(req.Offset),
			Max:    "+inf",
			Offset: 0,
			Count:  1000,
		}

		key = util.KeyMsgSync(req.Uin)
		cmd := rc.ZRangeByScore(ctx, key, &zr)
		val, errr := cmd.Result()
		if errr != nil {
			return errr
		}

		var keys []string
		for _, v := range val {
			key = util.KeyMsg(req.Uin, cast.ToInt64(v))
			keys = append(keys, key)
		}
		if len(keys) == 0 {
			break
		}

		// 同步库中存在，而消息库中却不存在
		// 发生这种情况是因为，消息库中的消息过期已从redis中清除了，但是同步库中的消息id还未即时跑批处理清理掉
		var dirtyMembers []interface{}
		rr, errr := rc.MGet(ctx, keys...).Result()
		for i, v := range rr {
			if v == nil {
				dirtyMembers = append(dirtyMembers, val[i])
				continue
			}
		}
		if len(dirtyMembers) == 0 {
			break
		} else {
			key = util.KeyMsgSync(req.Uin)
			if _, err = rc.ZRem(ctx, key, dirtyMembers...).Result(); err != nil {
				return
			}
		}
	}

	return
}

func (l *Chat) MsgAck(ctx context.Context, req *chat.MsgAckReq, rsp *chat.MsgAckRsp) (err error) {
	// TODO: 优化
	db := runtime.GetDB()
	msg := model.Msg{Id: req.Id}
	err = db.Take(&msg).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// TODO: ================
			//err = errno.ErrCustom("消息ID不存在")
		}
		log.Error(err.Error())
		return
	}
	rc := runtime.GetRedisClient()
	key := util.KeyMsgSync(req.Uin)
	rc.ZRemRangeByScore(ctx, key, "-inf", cast.ToString(msg.SendTime))

	//key = util.KeyConvMsgSync(req.Uin, msg.Target)
	//rc.ZRemRangeByScore(ctx, key, "-inf", cast.ToString(msg.SendTime))

	return
}

func (l *Chat) sendC2C(ctx context.Context, req *chat.SendReq, rsp *chat.SendRsp) (err error) {
	now := time.Now().UnixMilli()
	id := idgen.Next()
	p := common.Msg{
		Id:            id,
		ConvType:      req.ConvType,
		Type:          req.MsgType,
		Content:       req.Content,
		Sender:        req.Sender,
		Target:        req.Target,
		SendTime:      now,
		ClientUuid:    req.ClientUuid,
		AtUserList:    req.AtUserList,
		IsTransparent: req.IsTransparent,
	}

	if err = l.createConversation(ctx, req.Sender, req.Target, constant.ConvTypeC2C, &p); err != nil {
		return
	}

	if err = l.createConversation(ctx, req.Target, req.Sender, constant.ConvTypeC2C, &p); err != nil {
		return
	}

	// PUBLISH 两条
	p.Owner = req.Sender
	b, err := proto.Marshal(&p)
	if err != nil {
		return
	}
	m := &nats.Msg{
		Subject: "MSGS.new",
		Reply:   "",
		Data:    b,
		Sub:     nil,
	}
	js := runtime.GetJS()
	js.PublishMsg(m)
	p.Owner = req.Target
	b, err = proto.Marshal(&p)
	if err != nil {
		return
	}
	m.Data = b
	js.PublishMsg(m)

	// TODO: 移除code,message两个字段
	*rsp = chat.SendRsp{
		Code:       0,
		Message:    "",
		Id:         id,
		SendTime:   now,
		ClientUuid: req.ClientUuid,
	}

	return
}

func (l *Chat) sendC2G(ctx context.Context, req *chat.SendReq, rsp *chat.SendRsp) (err error) {
	now := time.Now().UnixMilli()

	id := idgen.Next()
	p := common.Msg{
		Id:            id,
		ConvType:      req.ConvType,
		Type:          req.MsgType,
		Content:       req.Content,
		Sender:        req.Sender,
		Target:        req.Target,
		SendTime:      now,
		ClientUuid:    req.ClientUuid,
		AtUserList:    req.AtUserList,
		IsTransparent: req.IsTransparent,
	}

	db := runtime.GetDB()
	var members []*model.GroupMember
	cond := model.GroupMember{GroupId: req.Target}
	if err = db.Where(&cond).Find(&members).Error; err != nil {
		return
	}

	js := runtime.GetJS()
	for _, v := range members {
		p.Owner = v.Member
		b, err := proto.Marshal(&p)
		if err != nil {
			continue
		}
		m := &nats.Msg{
			Subject: "MSGS.new",
			Reply:   "",
			Data:    b,
			Sub:     nil,
		}
		l.createConversation(ctx, v.Member, req.Target, constant.ConvTypeGroup, &p)

		js.PublishMsg(m)
	}

	//if err = l.createConversation(ctx, req.Sender, req.Target, constant.ConvTypeGroup, &p); err != nil {
	//	return
	//}

	*rsp = chat.SendRsp{
		Code:       0,
		Message:    "",
		Id:         id,
		SendTime:   now,
		ClientUuid: req.ClientUuid,
	}

	return
}

func (l *Chat) createConversation(ctx context.Context, owner, target string, convType int, m *common.Msg) (err error) {
	rc := runtime.GetRedisClient()
	rc.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		now := time.Now().UnixMilli()
		member := redis.Z{
			Score:  float64(now),
			Member: target,
		}

		key := util.KeyConvSync(owner)
		pipe.ZAdd(ctx, key, &member)
		pipe.Expire(ctx, key, time.Duration(constant.ConvKeepDays*24)*time.Hour)

		conv := &common.Conversation{
			Type:   convType,
			Target: target,
		}

		key = util.KeyConv(owner, target)
		v, err := pipe.Get(ctx, key).Result()
		if err != nil {
			if err != redis.Nil {
				return err
			}
		} else {
			json.Unmarshal([]byte(v), conv)
		}
		conv.LastMsg = m
		if m.Target == owner {
			conv.UnreadCount += 1
		}
		conv.UpdatedAt = now

		b, _ := json.Marshal(conv)
		pipe.SetEX(ctx, key, string(b), time.Duration(constant.ConvKeepDays*24)*time.Hour)
		return nil
	})
	return
}

type MsgRecall struct {
	Operator string `json:"operator"`
	Id       int64  `json:"id"`
}

func (l *Chat) Recall(ctx context.Context, req *chat.RecallReq, rsp *chat.RecallRsp) (err error) {
	rc := runtime.GetRedisClient()

	v := model.Msg{Sender: req.Uin, Id: req.Id}
	db := runtime.GetDB()
	if err = db.Take(&v).Error; err != nil {
		return
	}

	if time.Now().Unix()-cast.ToInt64(v.SendTime)/1000 > 120 {
		err = errors.New("不能撤回两分钟以前的消息")
		return
	}

	if v.ConvType == constant.ConvTypeC2C {
		rc.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			key1 := util.KeyMsg(v.Sender, v.Id)
			key2 := util.KeyMsg(v.Target, v.Id)
			pipe.Del(ctx, key1, key2)

			//key1 = util.KeyConvMsgSync(v.Sender, v.Target)
			//pipe.ZRem(ctx, key1, v.Id)
			//key2 = util.KeyConvMsgSync(v.Target, v.Sender)
			//pipe.ZRem(ctx, key2, v.Id)
			//var key string
			//if v.Sender < v.Target {
			//	key = util.KeyConvMsg(v.Sender, v.Target, v.Id)
			//} else {
			//	key = util.KeyConvMsg(v.Target, v.Sender, v.Id)
			//}
			//pipe.Del(ctx, key)
			return nil
		})
	} else if v.ConvType == constant.ConvTypeGroup {
		var members []*model.GroupMember
		cond := model.GroupMember{
			GroupId: v.Target,
		}
		if e := db.Where(&cond).Find(&members).Error; e != nil {
			log.Error(e.Error())
			err = e
			return
		}

		var keys []string
		for _, m := range members {
			key := util.KeyMsg(m.Member, v.Id)
			keys = append(keys, key)
		}
		rc.Del(ctx, keys...)

		//for _, m := range members {
		//	rc.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		//		key := util.KeyConvMsgSync(m.Member, v.Target)
		//		pipe.ZRem(ctx, key, v.Id)
		//		key = util.KeyConvMsg(m.Member, v.Target, v.Id)
		//		pipe.Del(ctx, key)
		//		return nil
		//	})
		//}

	}

	db.Delete(&v)

	m := MsgRecall{
		Operator: req.Uin,
		Id:       req.Id,
	}
	b, _ := json.Marshal(m)
	reqL := chat.SendReq{
		ConvType: v.ConvType,
		MsgType:  constant.MsgRecall,
		Sender:   v.Sender,
		Target:   v.Target,
		Content:  string(b),
	}
	rspL := chat.SendRsp{}
	l.SendMsg(ctx, &reqL, &rspL)

	return
}

func (l *Chat) DeleteMsg(ctx context.Context, req *chat.DeleteMsgReq, rsp *chat.DeleteMsgRsp) (err error) {
	if len(req.Ids) == 0 {
		return
	}

	rc := runtime.GetRedisClient()

	var members []interface{}
	for _, id := range req.Ids {
		members = append(members, id)
	}
	key := util.KeyMsgSync(req.Uin)
	if err = rc.ZRem(context.Background(), key, members...).Err(); err != nil {
		return
	}

	//rsp = &api.DeleteMsgRsp{}
	return
}
