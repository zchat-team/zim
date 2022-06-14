// Code generated by protoc-gen-zmicro-gin. DO NOT EDIT.
// versions:
// - protoc-gen-zmicro-gin v0.1.0
// - protoc                v3.19.0
// source: api/rest/group/group.proto

package group

import (
	context "context"
	errors "errors"
	gin "github.com/gin-gonic/gin"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = errors.New
var _ = context.TODO
var _ = gin.New

type GroupHTTPServer interface {
	Add(context.Context, *AddReq, *AddRsp) error
	Create(context.Context, *CreateReq, *CreateRsp) error
	Validate(context.Context, any) error
	ErrorEncoder(c *gin.Context, err error, isBadRequest bool)
}

type UnimplementedGroupHTTPServer struct{}

func (*UnimplementedGroupHTTPServer) Add(context.Context, *AddReq, *AddRsp) error {
	return errors.New("method Add not implemented")
}
func (*UnimplementedGroupHTTPServer) Create(context.Context, *CreateReq, *CreateRsp) error {
	return errors.New("method Create not implemented")
}
func (*UnimplementedGroupHTTPServer) Validate(context.Context, any) error { return nil }
func (*UnimplementedGroupHTTPServer) ErrorEncoder(c *gin.Context, err error, isBadRequest bool) {
	var code = 500
	if isBadRequest {
		code = 400
	}
	c.String(code, err.Error())
}

func RegisterGroupHTTPServer(g *gin.RouterGroup, srv GroupHTTPServer) {
	r := g.Group("")
	r.POST("/zim/groups", _Group_Create0_HTTP_Handler(srv))
	r.POST("/zim/groups/:group_id/members", _Group_Add0_HTTP_Handler(srv))
}

func _Group_Create0_HTTP_Handler(srv GroupHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		shouldBind := func(req any) error {
			if err := c.ShouldBind(req); err != nil {
				return err
			}
			return srv.Validate(c.Request.Context(), req)
		}

		var req CreateReq
		var rsp CreateRsp
		if err := shouldBind(&req); err != nil {
			srv.ErrorEncoder(c, err, true)
			return
		}
		err := srv.Create(c.Request.Context(), &req, &rsp)
		if err != nil {
			srv.ErrorEncoder(c, err, false)
			return
		}
		c.JSON(200, rsp)
	}
}

func _Group_Add0_HTTP_Handler(srv GroupHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		shouldBind := func(req any) error {
			if err := c.ShouldBind(req); err != nil {
				return err
			}
			if err := c.ShouldBindUri(req); err != nil {
				return err
			}
			return srv.Validate(c.Request.Context(), req)
		}

		var req AddReq
		var rsp AddRsp
		if err := shouldBind(&req); err != nil {
			srv.ErrorEncoder(c, err, true)
			return
		}
		err := srv.Add(c.Request.Context(), &req, &rsp)
		if err != nil {
			srv.ErrorEncoder(c, err, false)
			return
		}
		c.JSON(200, rsp)
	}
}
