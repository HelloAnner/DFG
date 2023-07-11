package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/logger/accesslog"
)

// go get github.com/hertz-contrib/logger/accesslog

// 访问日志可以收集所有 HTTP 请求的详细信息，包括时间、端口、请求方法等。Hertz 也提供了 access log 的 实现，这里的实现参考了 fiber。

// 默认格式
// [${time}] ${status} - ${latency} ${method} ${path}

// 支持的格式  https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/middleware/access-log/#支持的标签

func main() {
	h := server.Default(
		server.WithHostPorts(":18080"))

	h.Use(accesslog.New(accesslog.WithFormat("[${time}] ${status} - ${latency} ${method} ${path} ${queryParams}")))

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, utils.H{"msg": "pong"})
	})

	h.Spin()
}
