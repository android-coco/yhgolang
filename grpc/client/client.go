package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	myrpc "yhgolang/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := myrpc.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &myrpc.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
