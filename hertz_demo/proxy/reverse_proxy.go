package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/reverseproxy"
)

// 反向代理在计算机网络中是代理服务器的一种。
// 服务器根据客户端的请求，从其关系的一组或多组后端服务器（如 Web 服务器）上获取资源，
// 然后再将这些资源返回给客户端，客户端只会得知反向代理的 IP 地址，而不知道在代理服务器后面的服务器集群的存在。
// https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/proxy/
// 反向代理可以模拟服务端

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8000"))
	// 设置目标地址
	proxy, err := reverseproxy.NewSingleHostReverseProxy("http://127.0.0.1:8000/proxy")
	if err != nil {
		panic(err)
	}
	h.GET("/proxy/backend", func(cc context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{
			"msg": "proxy success!!",
		})
	})
	// 设置代理
	h.GET("/backend", proxy.ServeHTTP)
	h.Spin()
}
