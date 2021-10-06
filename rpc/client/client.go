package main

import (
	"context"
	hello_grpc "demo-for-golang/rpc/service"
	"fmt"
	"google.golang.org/grpc"
)

//
// @author Anner on 2021/10/6

func main() {
	conn,err := grpc.Dial(":8080",grpc.WithInsecure())
	if err!=nil{
		fmt.Printf("connect server fail : %s\n",err)
		return
	}

	defer conn.Close()

	// create client
	c:= hello_grpc.NewGreeterClient(conn)

	response,err:= c.SayHello(context.Background(),&hello_grpc.HelloRequest{Name: "Anner"})
	if err!=nil{
		fmt.Printf("call hello service fail: %s\n",err)
		return
	}

	fmt.Printf("response : %s\n",response.Message)
}