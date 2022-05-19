package main

import (
	"github.com/smallnest/rpcx/server"
	"github.com/zmicro-team/zim/app/chat/internal/service"
	"github.com/zmicro-team/zim/pkg/runtime"
	"github.com/zmicro-team/zmicro"
	"github.com/zmicro-team/zmicro/core/log"
)

func main() {
	app := zmicro.New(
		zmicro.InitRpcServer(InitRpcServer),
		zmicro.Before(func() error {
			runtime.Setup()
			return nil
		}),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func InitRpcServer(s *server.Server) error {
	if err := s.RegisterName("Chat", service.GetChatService(), ""); err != nil {
		log.Fatal(err.Error())
	}
	if err := s.RegisterName("Conv", service.GetConvService(), ""); err != nil {
		log.Fatal(err.Error())
	}
	return nil
}
