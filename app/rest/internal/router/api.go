package router

import (
	"github.com/gin-gonic/gin"

	"github.com/zmicro-team/zim/api/rest/chat"
	"github.com/zmicro-team/zim/api/rest/conv"
	"github.com/zmicro-team/zim/api/rest/group"
	"github.com/zmicro-team/zim/api/rest/user"

	sChat "github.com/zmicro-team/zim/app/rest/internal/service/chat"
	sConv "github.com/zmicro-team/zim/app/rest/internal/service/conv"
	sGroup "github.com/zmicro-team/zim/app/rest/internal/service/group"
	sUser "github.com/zmicro-team/zim/app/rest/internal/service/user"
)

func RegisterAPI(r *gin.Engine) {
	//Swagger(r)
	g := r.Group("/api")

	chat.RegisterImHTTPServer(g, sChat.GetService())
	conv.RegisterConvHTTPServer(g, sConv.GetService())
	group.RegisterGroupHTTPServer(g, sGroup.GetService())
	user.RegisterClientHTTPServer(g, sUser.GetService())
}
