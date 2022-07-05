//go:build !noswag
// +build !noswag

package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"

	"github.com/zchat-team/zchat-server/docs"
)

func Swagger(r gin.IRouter) {
	r.GET("/swagger/*any", SwaggerHandler())
}

func SwaggerHandler() gin.HandlerFunc {
	swag.Register(swag.Name, new(docs.Docs))
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
