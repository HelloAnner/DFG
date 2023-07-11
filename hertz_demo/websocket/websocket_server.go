package main

import (
	"context"
	"flag"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/websocket"
	"log"
)

// go get github.com/hertz-contrib/websocket

// Hertz 提供了 WebSocket 的支持，参考 gorilla/websocket 库使用 hijack 的方式在 Hertz 进行了适配，用法和参数基本保持一致

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.HertzUpgrader{} // use default options

func echo(_ context.Context, c *app.RequestContext) {
	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
}

func home(_ context.Context, c *app.RequestContext) {
	c.SetContentType("text/html; charset=utf-8")
	//homeTemplate.Execute(c, "ws://"+string(c.Host())+"/echo")
}

func main() {
	flag.Parse()
	h := server.Default(server.WithHostPorts(*addr))
	// https://github.com/cloudwego/hertz/issues/121
	h.NoHijackConnPool = true
	h.GET("/", home)
	h.GET("/echo", echo)
	h.Spin()
}

//var homeTemplate = ""
