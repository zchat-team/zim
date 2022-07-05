package router

import (
	"github.com/gin-gonic/gin"

	sChat "github.com/zchat-team/zim/app/rest/internal/service/chat"
	sConv "github.com/zchat-team/zim/app/rest/internal/service/conv"
	sGroup "github.com/zchat-team/zim/app/rest/internal/service/group"
	sUser "github.com/zchat-team/zim/app/rest/internal/service/user"
	"github.com/zchat-team/zim/proto/http/rest/chat"
	"github.com/zchat-team/zim/proto/http/rest/conv"
	"github.com/zchat-team/zim/proto/http/rest/group"
	"github.com/zchat-team/zim/proto/http/rest/user"
)

func RegisterAPI(r *gin.Engine) {
	Swagger(r)
	g := r.Group("/api")

	chat.RegisterImHTTPServer(g, sChat.GetService())
	conv.RegisterConvHTTPServer(g, sConv.GetService())
	group.RegisterGroupHTTPServer(g, sGroup.GetService())
	user.RegisterClientHTTPServer(g, sUser.GetService())
}
