package main

import (
	"context"
	"crypto/tls"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol"
)

// 正向代理是一种特殊的网络服务，允许一个网络终端（一般为客户端）通过这个服务与另一个网络终端（一般为服务器）进行非直接的连接。
// 一些网关、路由器等网络设备具备网络代理功能。一般认为代理服务有利于保障网络终端的隐私或安全，防止攻击。
// https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/proxy/
// 正向代理的例子一般都是客户端发出请求

func main() {
	proxyURL := "http://127.0.0.1:7890"

	// 将代理的 uri 转成 *protocol.URI 的形式
	parsedProxyURL := protocol.ParseURI(proxyURL)

	clientCfg := &tls.Config{
		InsecureSkipVerify: true,
	}

	c, err := client.NewClient(client.WithTLSConfig(clientCfg), client.WithDialer(standard.NewDialer()))
	if err != nil {
		return
	}

	// 设置代理
	c.SetProxy(protocol.ProxyURI(parsedProxyURL))

	upstreamURL := "https://www.google.com"

	_, body, _ := c.Get(context.Background(), nil, upstreamURL)

	println(string(body))
}
