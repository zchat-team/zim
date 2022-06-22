package main

import (
	"github.com/smallnest/rpcx/server"
	"github.com/zchat-team/zim/app/chat/internal/model"
	"github.com/zchat-team/zim/app/chat/internal/service"
	"github.com/zchat-team/zim/pkg/runtime"
	"github.com/zmicro-team/zmicro"
	"github.com/zmicro-team/zmicro/core/log"
)

func main() {
	app := zmicro.New(
		zmicro.InitRpcServer(InitRpcServer),
		zmicro.Before(before),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitRpcServer(s *server.Server) error {
	if err := s.RegisterName("Chat", service.GetChatService(), ""); err != nil {
		log.Fatal(err)
	}
	if err := s.RegisterName("Conv", service.GetConvService(), ""); err != nil {
		log.Fatal(err)
	}
	if err := s.RegisterName("Group", service.GetGroupService(), ""); err != nil {
		log.Fatal(err)
	}
	return nil
}

func before() error {
	runtime.Setup()
	db := runtime.GetDB()
	if err := db.AutoMigrate(
		&model.Msg{},
		&model.User{},
		&model.Group{},
		&model.GroupMember{},
	); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
