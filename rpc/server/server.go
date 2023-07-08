package main

import (
	"context"
	hello_grpc "dfg/rpc/service"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

//
// @author Anner on 2021/10/6

type server struct {
}

func (s *server) SayHello(ctx context.Context, request *hello_grpc.HelloRequest) (*hello_grpc.HelloReply, error) {
	return &hello_grpc.HelloReply{
		Message: "hello" + request.GetName(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("failed to list port: %s\n", err)
		return
	}

	// rpc server
	s := grpc.NewServer()
	// register service
	hello_grpc.RegisterGreeterServer(s, &server{})

	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("start serve fail : %s\n", err)
		return
	}
}
