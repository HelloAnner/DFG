package main

import (
	"context"
	"dfg/kitex_demo/hello/kitex_gen/api"
	"dfg/kitex_demo/hello/kitex_gen/api/hello"
	"github.com/cloudwego/kitex/client"
	"log"
	"time"
)

func main() {
	client, err := hello.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	a, b := 1, 1

	for {
		req := &api.Request{Message: "my request"}
		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)

		addReq := &api.AddRequest{
			First:  int64(a),
			Second: int64(b),
		}

		addResp, err := client.Add(context.Background(), addReq)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(addResp)

		a = b
		b = int(addResp.Sum)

		time.Sleep(time.Second)
	}
}
