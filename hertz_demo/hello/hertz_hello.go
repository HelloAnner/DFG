package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// https://github.com/cloudwego/hertz-examples/tree/main/hello
func main() {
	// server.Default() creates a Hertz with recovery middleware.
	// If you need a pure hertz, you can use server.New()
	h := server.Default()

	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "Hello hertz!")
	})

	h.OnRun = append(h.OnRun, func(ctx context.Context) error {
		fmt.Printf("server is online now\n")
		return nil
	})

	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		println("shutdown now")
	})

	hlog.Info("This is  Hertz Log")

	h.Spin()
}
