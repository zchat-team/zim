// Code generated by protoc-gen-zmicro-gin. DO NOT EDIT.
// versions:
// - protoc-gen-zmicro-gin v0.1.0
// - protoc                v3.19.0
// source: proto/http/demo/user/user.proto

package user

import (
	context "context"
	errors "errors"
	gin "github.com/gin-gonic/gin"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = errors.New
var _ = context.TODO
var _ = gin.New

type UserHTTPServer interface {
	Get(context.Context, *GetReq, *GetRsp) error
	MGet(context.Context, *MGetReq, *MGetRsp) error
	Search(context.Context, *SearchReq, *SearchRsp) error
	Validate(context.Context, any) error
	ErrorEncoder(c *gin.Context, err error, isBadRequest bool)
}

type UnimplementedUserHTTPServer struct{}

func (*UnimplementedUserHTTPServer) Get(context.Context, *GetReq, *GetRsp) error {
	return errors.New("method Get not implemented")
}
func (*UnimplementedUserHTTPServer) MGet(context.Context, *MGetReq, *MGetRsp) error {
	return errors.New("method MGet not implemented")
}
func (*UnimplementedUserHTTPServer) Search(context.Context, *SearchReq, *SearchRsp) error {
	return errors.New("method Search not implemented")
}
func (*UnimplementedUserHTTPServer) Validate(context.Context, any) error { return nil }
func (*UnimplementedUserHTTPServer) ErrorEncoder(c *gin.Context, err error, isBadRequest bool) {
	var code = 500
	if isBadRequest {
		code = 400
	}
	c.String(code, err.Error())
}

func RegisterUserHTTPServer(g *gin.RouterGroup, srv UserHTTPServer) {
	r := g.Group("")
	r.POST("/user/search", _User_Search0_HTTP_Handler(srv))
	r.POST("/user/get", _User_Get0_HTTP_Handler(srv))
	r.POST("/user/mget", _User_MGet0_HTTP_Handler(srv))
}

func _User_Search0_HTTP_Handler(srv UserHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		shouldBind := func(req any) error {
			if err := c.ShouldBind(req); err != nil {
				return err
			}
			return srv.Validate(c.Request.Context(), req)
		}

		var req SearchReq
		var rsp SearchRsp
		if err := shouldBind(&req); err != nil {
			srv.ErrorEncoder(c, err, true)
			return
		}
		err := srv.Search(c.Request.Context(), &req, &rsp)
		if err != nil {
			srv.ErrorEncoder(c, err, false)
			return
		}
		c.JSON(200, rsp)
	}
}

func _User_Get0_HTTP_Handler(srv UserHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		shouldBind := func(req any) error {
			if err := c.ShouldBind(req); err != nil {
				return err
			}
			return srv.Validate(c.Request.Context(), req)
		}

		var req GetReq
		var rsp GetRsp
		if err := shouldBind(&req); err != nil {
			srv.ErrorEncoder(c, err, true)
			return
		}
		err := srv.Get(c.Request.Context(), &req, &rsp)
		if err != nil {
			srv.ErrorEncoder(c, err, false)
			return
		}
		c.JSON(200, rsp)
	}
}

func _User_MGet0_HTTP_Handler(srv UserHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		shouldBind := func(req any) error {
			if err := c.ShouldBind(req); err != nil {
				return err
			}
			return srv.Validate(c.Request.Context(), req)
		}

		var req MGetReq
		var rsp MGetRsp
		if err := shouldBind(&req); err != nil {
			srv.ErrorEncoder(c, err, true)
			return
		}
		err := srv.MGet(c.Request.Context(), &req, &rsp)
		if err != nil {
			srv.ErrorEncoder(c, err, false)
			return
		}
		c.JSON(200, rsp)
	}
}