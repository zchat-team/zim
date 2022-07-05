//go:build noswag
// +build noswag

package router

import "github.com/gin-gonic/gin"

func Swagger(gin.IRouter) {}
