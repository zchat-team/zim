package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zchat-team/zim/pkg/util"
	"gorm.io/gorm"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/zchat-team/zim/app/task/internal/client"
	"github.com/zchat-team/zim/app/task/internal/model"
	"github.com/zchat-team/zim/pkg/constant"
	"github.com/zchat-team/zim/pkg/runtime"
	"github.com/zchat-team/zim/proto/common"
	"github.com/zchat-team/zim/proto/sess"
	"github.com/zmicro-team/zmicro/core/log"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	go s.consumeMsg()

	log.Info("Dispatch Server Started")

	return nil
}

func (s *Server) Stop() error {
	return nil
}

func (s *Server) consumeMsg() {
	js := runtime.GetJS()
	sub, err := js.PullSubscribe("MSGS.received", "TASK")
	if err != nil {
		log.Fatal(err)
	}

	for {
		msgs, err := sub.Fetch(10)
		if err != nil {
			if errors.Is(err, nats.ErrTimeout) {
				continue
			}
			log.Error(err.Error())
		} else {
			for _, m := range msgs {
				msg := common.Msg{}
				if err := proto.Unmarshal(m.Data, &msg); err != nil {
					m.Ack()
					continue
				}

				if err := s.onMsg(&msg); err == nil {
					m.Ack()
				}
			}
		}
	}
}

func (s *Server) onMsg(m *common.Msg) error {
	if err := s.storeRedis(m); err != nil {
		return err
	}

	s.push(m)
	s.storeMysql(m)

	return nil
}

func (s *Server) storeRedis(m *common.Msg) error {
	// TODO: 判断透传消息，不存储

	member := redis.Z{
		Score:  float64(m.SendTime),
		Member: m.Id,
	}

	rc := runtime.GetRedisClient()
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	// TODO: context
	ctx := context.Background()
	if _, err := rc.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		key := util.KeyMsgSync(m.Owner)
		pipe.ZAdd(ctx, key, &member)
		pipe.Expire(ctx, key, time.Duration(constant.MsgKeepDays*24)*time.Hour)

		key = util.KeyMsg(m.Owner, m.Id)
		pipe.SetEX(ctx, key, string(b), time.Duration(constant.MsgKeepDays*24)*time.Hour)

		// TODO: 方案二，优化或者直接废弃
		if m.ConvType == constant.ConvTypeC2C {
			if m.Owner == m.Sender {
				key = util.KeyConvMsgSync(m.Owner, m.Target)
			} else {
				key = util.KeyConvMsgSync(m.Owner, m.Sender)
			}
			pipe.ZAdd(ctx, key, &member)

			pipe.Expire(ctx, key, time.Duration(constant.MsgKeepDays*24)*time.Hour)

			if m.Sender < m.Target {
				key = util.KeyConvMsg(m.Sender, m.Target, m.Id)
			} else {
				key = util.KeyConvMsg(m.Target, m.Sender, m.Id)
			}
			pipe.SetEX(ctx, key, string(b), time.Duration(constant.MsgKeepDays*24)*time.Hour)

		} else {
			key = util.KeyConvMsgSync(m.Owner, m.Target)
			pipe.ZAdd(ctx, key, &member)
			pipe.Expire(ctx, key, time.Duration(constant.MsgKeepDays*24)*time.Hour)
			key = util.KeyConvMsg(m.Owner, m.Target, m.Id)
			pipe.SetEX(ctx, key, string(b), time.Duration(constant.MsgKeepDays*24)*time.Hour)
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (s *Server) push(m *common.Msg) {
	// 获取在线状态
	sessClient := client.GetSessClient()
	if sessClient != nil {
		log.Infof("Uin=%s", m.Owner)
		req := sess.GetOnlineReq{Uin: m.Owner}
		rsp, err := sessClient.GetOnline(context.Background(), &req)
		if err != nil {
			log.Error(err)
			return
		}

		b, err := proto.Marshal(m)
		if err != nil {
			log.Error(err)
			return
		}

		nc := runtime.GetNC()
		for _, v := range rsp.Devices {
			// online
			if v.Status == 1 {
				// 在线推送
				var onlineDevices []string
				onlineDevices = append(onlineDevices, v.DeviceId)
				pushMsg := common.PushMsg{
					Server:  v.Server,
					Devices: onlineDevices,
					Msg:     b,
				}
				bb, err := proto.Marshal(&pushMsg)
				if err != nil {
					log.Error(err)
					continue
				}

				mm := nats.Msg{
					Subject: fmt.Sprintf("push.online.%s", v.Server),
					Reply:   "",
					Header:  nil,
					Data:    bb,
					Sub:     nil,
				}
				if err := nc.PublishMsg(&mm); err != nil {
					log.Error(err)
				}
			} else if v.Status == 2 {
				// TODO: 离线推送
			}
		}
	} else {
		log.Info("client is null")
	}
}

func (s *Server) storeMysql(m *common.Msg) {
	var atUserList string
	if len(m.AtUserList) > 0 {
		b, _ := json.Marshal(m.AtUserList)
		atUserList = string(b)
	}

	db := runtime.GetDB()
	msg := model.Msg{
		Id:         m.Id,
		ConvType:   int(m.ConvType),
		Content:    m.Content,
		Type:       int(m.Type),
		DeletedAt:  0,
		Sender:     m.Sender,
		Target:     m.Target,
		AtUserList: atUserList,
		ReadTime:   0,
		SendTime:   m.SendTime,
		ClientUuid: m.ClientUuid,
	}

	if err := db.Take(&model.Msg{Id: m.Id}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
			return
		}
	} else {
		return
	}

	if err := db.Create(&msg).Error; err != nil {
		log.Error(err)
	}
}
