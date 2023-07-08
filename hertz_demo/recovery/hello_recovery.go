package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/middleware/recovery/

// Recovery 中间件是 Hertz 框架预置的中间件，使用 server.Default() 可以默认注册该中间件，为 Hertz 框架提供 panic 恢复的功能。

// Recovery 中间件会恢复 Hertz 框架运行中的任何 panic，在 panic 发生之后，Recover 中间件会默认打印出 panic 的时间、内容和堆栈信息
// 同时通过*app.RequestContext将返回响应的状态码设置成 500。
func main() {
	h := server.New()
	h.Use(recovery.Recovery(recovery.WithRecoveryHandler(CustomRecoveryHandler)))

	h.GET("/test", func(ctx context.Context, c *app.RequestContext) {
		panic("test")
	})

	h.Spin()
}

func CustomRecoveryHandler(c context.Context, ctx *app.RequestContext, err interface{}, stack []byte) {
	hlog.SystemLogger().CtxErrorf(c, "[Recovery] err=%v\nstack=%s", err, stack)
	hlog.SystemLogger().Infof("Client: %s", ctx.Request.Header.UserAgent())
	ctx.AbortWithStatus(consts.StatusInternalServerError)
}
