package main

import (
	"github.com/zmicro-team/zim/app/task/internal/app"
	"github.com/zmicro-team/zmicro/core/log"
)

func main() {
	a := app.New()
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
