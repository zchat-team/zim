package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zchat-team/zim/pkg/validate"
	zgin "github.com/zmicro-team/zmicro/core/transport/http"
	"net/http"
	"time"
)

func Setup(engine *gin.Engine) {
	gin.DisableBindValidation()
	validate.RegisterValidation(zgin.Validator())

	engine.NoMethod(func(ctx *gin.Context) {
		ctx.String(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	})

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	})

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST, OPTIONS, GET, PUT, PATCH, DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	RegisterAPI(engine)
}
