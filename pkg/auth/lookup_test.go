package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func TestNewLookupIngoreInvalidLookupPair(t *testing.T) {
	NewLookup("header:Authorization:Bearer,xxx")
}

func TestLookupHeader(t *testing.T) {
	lk := NewLookup("header:Authorization:Bearer")

	t.Run("miss header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/get", nil)

		srv := gin.New()
		srv.GET("/get", func(c *gin.Context) {
			token, err := lk.GetToken(c)
			require.Error(t, err)
			require.Empty(t, token)
		})
		srv.ServeHTTP(httptest.NewRecorder(), req)
	})
	t.Run("from header with Bearer", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/get", nil)
		req.Header.Add("Authorization", "Bearer xxxxxx")

		srv := gin.New()
		srv.GET("/get", func(c *gin.Context) {
			token, err := lk.GetToken(c)
			require.NoError(t, err)
			require.Equal(t, "xxxxxx", token)
		})
		srv.ServeHTTP(httptest.NewRecorder(), req)
	})
	t.Run("from header but empty", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/get", nil)
		req.Header.Add("Authorization", "Bearer  ")

		srv := gin.New()
		srv.GET("/get", func(c *gin.Context) {
			token, err := lk.GetToken(c)
			require.Error(t, err)
			require.Empty(t, token)
		})
		srv.ServeHTTP(httptest.NewRecorder(), req)
	})
	t.Run("from header without Bearer", func(t *testing.T) {
		lk1 := NewLookup("header:Authorization")

		req := httptest.NewRequest(http.MethodGet, "/get", nil)
		req.Header.Add("Authorization", "xxxxxx")

		srv := gin.New()
		srv.GET("/get", func(c *gin.Context) {
			token, err := lk1.GetToken(c)
			require.NoError(t, err)
			require.Equal(t, "xxxxxx", token)
		})
		srv.ServeHTTP(httptest.NewRecorder(), req)
	})
	t.Run("from header but invalid value", func(t *testing.T) {
		lk1 := NewLookup("header:Authorization:Bearer")

		req := httptest.NewRequest(http.MethodGet, "/get", nil)
		req.Header.Add("Authorization", "xxxxxx")

		srv := gin.New()
		srv.GET("/get", func(c *gin.Context) {
			token, err := lk1.GetToken(c)
			require.Error(t, err)
			require.Equal(t, "", token)
		})
		srv.ServeHTTP(httptest.NewRecorder(), req)
	})
}

func TestLookupQuery(t *testing.T) {
	lk := NewLookup("query:token")

	t.Run("miss query", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/get", nil)

		srv := gin.New()
		srv.GET("/get", func(c *gin.Context) {
			token, err := lk.GetToken(c)
			require.Error(t, err)
			require.Empty(t, token)
		})
		srv.ServeHTTP(httptest.NewRecorder(), req)
	})
	t.Run("from query", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/get?token=xxxxxx", nil)

		srv := gin.New()
		srv.GET("/get", func(c *gin.Context) {
			token, err := lk.GetToken(c)
			require.NoError(t, err)
			require.Equal(t, "xxxxxx", token)
		})
		srv.ServeHTTP(httptest.NewRecorder(), req)
	})
}

func TestLookupCookie(t *testing.T) {
	lk := NewLookup("cookie:token")

	t.Run("miss cookie", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/get", nil)

		srv := gin.New()
		srv.GET("/get", func(c *gin.Context) {
			token, err := lk.GetToken(c)
			require.Error(t, err)
			require.Empty(t, token)
		})
		srv.ServeHTTP(httptest.NewRecorder(), req)
	})
	t.Run("from cookie", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/get", nil)
		req.AddCookie(&http.Cookie{
			Name:  "token",
			Value: "xxxxxx",
		})

		srv := gin.New()
		srv.GET("/get", func(c *gin.Context) {
			token, err := lk.GetToken(c)
			require.NoError(t, err)
			require.Equal(t, "xxxxxx", token)
		})
		srv.ServeHTTP(httptest.NewRecorder(), req)
	})
}

func TestLookupParam(t *testing.T) {
	lk := NewLookup("param:token")

	t.Run("miss param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/get/", nil)

		srv := gin.New()
		srv.GET("/get/:token", func(c *gin.Context) {
			token, err := lk.GetToken(c)
			require.Error(t, err)
			require.Empty(t, token)
		})
		srv.ServeHTTP(httptest.NewRecorder(), req)
	})
	t.Run("from param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/get/xxxxxx", nil)

		srv := gin.New()
		srv.GET("/get/:token", func(c *gin.Context) {
			token, err := lk.GetToken(c)
			require.NoError(t, err)
			require.Equal(t, "xxxxxx", token)
		})
		srv.ServeHTTP(httptest.NewRecorder(), req)
	})
}

func TestLookupMultiWay(t *testing.T) {
	lk := NewLookup("header:Authorization:Bearer,query:token,cookie:token,param:token")

	t.Run("from query", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/get?token=xxxxxx", nil)

		srv := gin.New()
		srv.GET("/get", func(c *gin.Context) {
			token, err := lk.GetToken(c)
			require.NoError(t, err)
			require.Equal(t, "xxxxxx", token)
		})
		srv.ServeHTTP(httptest.NewRecorder(), req)
	})
	t.Run("from header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/get", nil)
		req.Header.Add("Authorization", "Bearer xxxxxx")

		srv := gin.New()
		srv.GET("/get", func(c *gin.Context) {
			token, err := lk.GetToken(c)
			require.NoError(t, err)
			require.Equal(t, "xxxxxx", token)
		})
		srv.ServeHTTP(httptest.NewRecorder(), req)
	})
}
