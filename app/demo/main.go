package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zchat-team/zim/app/demo/internal/model"
	"github.com/zchat-team/zim/app/demo/internal/router"
	"github.com/zchat-team/zim/pkg/runtime"
	"github.com/zmicro-team/zmicro"
	"github.com/zmicro-team/zmicro/core/log"
)

func main() {
	app := zmicro.New(
		zmicro.Before(before),
		zmicro.InitHttpServer(InitHttpServer),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitHttpServer(r *gin.Engine) error {
	router.Setup(r)
	return nil
}

func before() error {
	runtime.Setup()
	db := runtime.GetDB()
	if err := db.AutoMigrate(
		&model.User{},
		&model.UserLoginLog{},
		&model.Friend{},
		&model.Application{},
		&model.Group{},
		&model.GroupMember{},
	); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
