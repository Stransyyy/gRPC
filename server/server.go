package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gRPC "github.com/Stransyyy/gRPC"
)

type server struct {
	gRPC.UnimplementedGreeterServer
}

func (*server) SayHello(ctx context.Context, req *gRPC.HelloRequest) (*gRPC.HelloReply, error) {
	if req.GetName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name cannot be empty")
	}

	fmt.Printf("SayHello function was invoked with %v\n", req)
	return &gRPC.HelloReply{
		Message: "Hello " + req.GetName(),
	}, nil
}

func main() {
	fmt.Println("Hello gRPC server")

	port := fmt.Sprintf(":%d", 8080)

	// Create a listener on TCP port 8080
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC server
	s := grpc.NewServer()

	// Register the Hello service with the server
	gRPC.RegisterGreeterServer(s, &server{})

	fmt.Printf("Server is running on port %s\n", port)

	// Start the gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
