package main

import (
	"github.com/zchat-team/zim/app/conn/internal/app"
	"github.com/zchat-team/zim/pkg/runtime"
	"github.com/zmicro-team/zmicro/core/log"
)

func main() {
	a := app.New(app.Before(func() error {
		runtime.Setup()
		return nil
	}))

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
