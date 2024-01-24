package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gRPC "github.com/Stransyyy/gRPC"
)

func main() {

	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	client := gRPC.NewGreeterClient(conn)

	response, err := client.SayHello(context.Background(), &gRPC.HelloRequest{Name: "Stransyyy"})
	if err != nil {
		log.Fatalf("The request failed: %v", err)
	}

	fmt.Printf("Response from server: %s", response.GetMessage())

}
