package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	"time"
)

// https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/middleware/cors/

// go get github.com/hertz-contrib/cors

// 对于跨源访问来说，如果是简单请求，本质上就是在 HTTP 请求头信息中添加一个 Origin 字段，用于描述本次请求来自哪个源，服务端可以直接响应

// 而对于非简单跨源访问请求来说（比如请求方法是 PUT 或 PATCH，Content-Type 字段类型是 application/json 等），会在正式通信之前，
// 发送一次 HTTP 预检请求（preflight），用于校验客户端是否有跨源资源访问权限，预检请求使用的方法是 OPTIONS，且这是浏览器自发的行为。
func main() {
	h := server.Default()
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	h.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		// 用于设置预检请求的有效期（有效期内不会发起重复的预检请求）
		MaxAge: 12 * time.Hour,
	}))
	h.Spin()
}
