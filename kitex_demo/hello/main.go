package main

import (
	api "dfg/kitex_demo/hello/kitex_gen/api/hello"
	"log"
)

// https://www.cloudwego.io/zh/docs/kitex/getting-started/

func main() {
	svr := api.NewServer(new(HelloImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
