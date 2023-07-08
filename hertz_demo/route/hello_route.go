package hertz_demo

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/basic_auth"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

//https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/route/

func routeDemo() {
	h := server.Default(server.WithHostPorts("127.0.0.1:19090"))

	h.GET("/get", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "get")
	})

	// 路由组
	v1 := h.Group("/v1")
	v1.GET("/get", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "for v1 get")
	})

	// 路由组 + 中间件
	v2 := h.Group("/v2", basic_auth.BasicAuth(map[string]string{"test": "test"}))
	v2.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "pong")
	})

	v3 := h.Group("/v3")

	v3.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	// 命名参数路由
	v3.GET("/hello/:name", func(c context.Context, ctx *app.RequestContext) {
		name := ctx.Param("name")
		ctx.String(consts.StatusOK, fmt.Sprintf("hello %s", name))
	})

	// 通用参数路由
	v3.GET("/any/*action", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "any thing")
	})

	hlog.Info(h.Routes())

	h.Spin()
}
