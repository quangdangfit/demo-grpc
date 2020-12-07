package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "demo-grpc/pb/user"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)
	request := &pb.GetRequest{
		Keyword: "params request",
	}
	resp, err := client.GetUsers(context.Background(), request)
	// To do something with resp from instance server response

	fmt.Println(resp)
}
