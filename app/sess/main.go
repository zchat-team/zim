package main

import (
	"github.com/smallnest/rpcx/server"
	"github.com/zmicro-team/zmicro"
	"github.com/zmicro-team/zmicro/core/log"

	"github.com/zchat-team/zim/app/sess/internal/service"
	"github.com/zchat-team/zim/pkg/runtime"
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
		log.Fatal(err)
	}
}

func InitRpcServer(s *server.Server) error {
	if err := s.RegisterName("Sess", service.GetService(), ""); err != nil {
		log.Fatal(err)
	}
	return nil
}
