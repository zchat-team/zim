package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/zmicro-team/zmicro/core/log"

	"github.com/zchat-team/zim/app/task/internal/client"
	"github.com/zchat-team/zim/app/task/internal/model"
	"github.com/zchat-team/zim/pkg/constant"
	"github.com/zchat-team/zim/pkg/runtime"
	"github.com/zchat-team/zim/pkg/util"
	"github.com/zchat-team/zim/proto/rpc/common"
	"github.com/zchat-team/zim/proto/rpc/sess"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	go s.consumeNew()
	go s.consumeTodo()

	log.Info("Dispatch Server Started")

	return nil
}

func (s *Server) Stop() error {
	return nil
}

func (s *Server) consumeTodo() {
	js := runtime.GetJS()
	sub, err := js.PullSubscribe("MSGS.todo", "TASK_TODO")
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

				if err := s.onTodo(&msg); err == nil {
					m.Ack()
				}
			}
		}
	}
}

func (s *Server) onTodo(m *common.Msg) error {
	if err := s.storeRedis(m); err != nil {
		return err
	}

	s.push(m)

	return nil
}

func (s *Server) storeRedis(m *common.Msg) error {
	if m.IsTransparent {
		return nil
	}

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

		// TODO: ????????????????????????????????????
		//if m.ConvType == constant.ConvTypeC2C {
		//	if m.Owner == m.Sender {
		//		key = util.KeyConvMsgSync(m.Owner, m.Target)
		//	} else {
		//		key = util.KeyConvMsgSync(m.Owner, m.Sender)
		//	}
		//	pipe.ZAdd(ctx, key, &member)
		//
		//	pipe.Expire(ctx, key, time.Duration(constant.MsgKeepDays*24)*time.Hour)
		//
		//	if m.Sender < m.Target {
		//		key = util.KeyConvMsg(m.Sender, m.Target, m.Id)
		//	} else {
		//		key = util.KeyConvMsg(m.Target, m.Sender, m.Id)
		//	}
		//	pipe.SetEX(ctx, key, string(b), time.Duration(constant.MsgKeepDays*24)*time.Hour)
		//
		//} else {
		//	key = util.KeyConvMsgSync(m.Owner, m.Target)
		//	pipe.ZAdd(ctx, key, &member)
		//	pipe.Expire(ctx, key, time.Duration(constant.MsgKeepDays*24)*time.Hour)
		//	key = util.KeyConvMsg(m.Owner, m.Target, m.Id)
		//	pipe.SetEX(ctx, key, string(b), time.Duration(constant.MsgKeepDays*24)*time.Hour)
		//}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (s *Server) push(m *common.Msg) {
	// ??????????????????
	sessClient := client.GetSessClient()
	if sessClient != nil {
		log.Infof("Uin=%s", m.Owner)
		req := sess.GetOnlineReq{Uin: m.Owner}
		rsp, err := sessClient.GetOnline(context.Background(), &req)
		if err != nil {
			log.Error(err)
			return
		}

		m.Owner = ""
		b, err := proto.Marshal(m)
		if err != nil {
			log.Error(err)
			return
		}

		nc := runtime.GetNC()
		for _, v := range rsp.Conns {
			// online
			if v.Status == 1 {
				// ????????????
				var onlines []string
				onlines = append(onlines, v.ConnId)
				pushMsg := common.PushMsg{
					Server: v.Server,
					Conns:  onlines,
					Msg:    b,
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
				// TODO: ????????????
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
		Sender:     m.Sender,
		Target:     m.Target,
		AtUserList: atUserList,
		ReadTime:   0,
		SendTime:   m.SendTime,
		ClientUuid: m.ClientUuid,
	}

	if err := db.Create(&msg).Error; err != nil {
		log.Error(err)
	}
}

func (s *Server) consumeNew() {
	js := runtime.GetJS()
	sub, err := js.PullSubscribe("MSGS.new", "TASK_NEW")
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
					log.Error(err)
					m.Ack()
					continue
				}

				if err := s.onNew(&msg); err == nil {
					m.Ack()
				}
			}
		}
	}
}

func (s *Server) onNew(m *common.Msg) (err error) {
	if m.ConvType == constant.ConvTypeC2C {
		err = s.onC2CMsg(m)
	} else if m.ConvType == constant.ConvTypeGroup {
		err = s.onGroupMsg(m)
	}

	if err != nil {
		return
	}
	// ???????????????????????????????????? MSGS.persist??????????????????????????????
	go func() {
		s.storeMysql(m)
	}()

	return
}

func (s *Server) onC2CMsg(m *common.Msg) error {
	js := runtime.GetJS()
	if m.Sender != "" {
		m.Owner = m.Sender
		b, err := proto.Marshal(m)
		if err != nil {
			return err
		}
		nm := &nats.Msg{
			Subject: "MSGS.todo",
			Reply:   "",
			Data:    b,
			Sub:     nil,
		}
		js.PublishMsg(nm)
	}

	if m.Target != "" {
		m.Owner = m.Target
		b, err := proto.Marshal(m)
		if err != nil {
			return err
		}
		nm := &nats.Msg{
			Subject: "MSGS.todo",
			Reply:   "",
			Data:    b,
			Sub:     nil,
		}
		js.PublishMsg(nm)
	}

	return nil
}

func (s *Server) onGroupMsg(m *common.Msg) (err error) {
	db := runtime.GetDB()
	var members []*model.GroupMember
	cond := model.GroupMember{GroupId: m.Target}
	if err = db.Where(&cond).Find(&members).Error; err != nil {
		return
	}

	js := runtime.GetJS()
	for _, v := range members {
		if v.Member == "" {
			continue
		}
		m.Owner = v.Member
		b, err := proto.Marshal(m)
		if err != nil {
			continue
		}
		nm := &nats.Msg{
			Subject: "MSGS.todo",
			Reply:   "",
			Data:    b,
			Sub:     nil,
		}

		js.PublishMsg(nm)
	}

	return
}
