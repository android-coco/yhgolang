package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	myrpc "yhgolang/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	myrpc.RegisterHelloServiceServer(grpcServer, &myrpc.HelloServiceImpl{})

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
