package main

import (
	"github.com/zmicro-team/zim/app/task/internal/app"
	"github.com/zmicro-team/zim/pkg/runtime"
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
