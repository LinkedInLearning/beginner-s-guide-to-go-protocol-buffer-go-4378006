package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/LinkedInLearning/beginner-s-guide-to-go-Proto-protocol-buffer-go-4378006/go/customer"
)

func main() {
	// a := customer.App{}
	// a.Initialize()
	// a.Run()

	lis, err := net.Listen("tcp", "localhost:9006")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCustomerServiceServer(grpcServer, newServer())
	fmt.Println("gRPC Server started and listening on port :9006")
	grpcServer.Serve(lis)
} 