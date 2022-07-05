package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	zgin "github.com/zmicro-team/zmicro/core/transport/http"

	"github.com/zchat-team/zim/app/demo/internal/service/passport"
	"github.com/zchat-team/zim/errno"
	"github.com/zchat-team/zim/pkg/auth"
	"github.com/zchat-team/zim/pkg/zcontext"
)

func authToken(token string) (acc *auth.Account, err error) {
	acc, err = passport.GetService().AuthToken(token)
	if err != nil {
		return
	}

	return
}

func CheckLogin(excludePrefixes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if checkPrefix(c.Request.URL.Path, excludePrefixes...) {
			// 不需要验证登录状态
			c.Next()
			return
		}

		token := zcontext.GetToken(c.Request.Context()) // 注意，不能传c,应该传c.Request.Context()
		acc, err := authToken(token)
		if err != nil {
			zgin.Error(c, errno.ErrInvalidToken())
			return
		}

		c.Request = c.Request.WithContext(zcontext.ContextWithAccount(c.Request.Context(), acc))
		c.Next()
	}
}

func checkPrefix(s string, prefixes ...string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
