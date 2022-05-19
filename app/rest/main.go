package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zmicro-team/zim/app/rest/internal/router"
	"github.com/zmicro-team/zim/pkg/runtime"
	"github.com/zmicro-team/zmicro"
	"github.com/zmicro-team/zmicro/core/log"
)

func main() {
	app := zmicro.New(zmicro.InitHttpServer(InitHttpServer))

	runtime.Setup()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitHttpServer(r *gin.Engine) error {
	router.Setup(r)
	return nil
}
