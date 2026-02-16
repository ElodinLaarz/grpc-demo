package main

import (
	"log"
	"net"

	pb "github.com/ElodinLaarz/grpc-demo/proto"
	srv "github.com/ElodinLaarz/grpc-demo/internal/server"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	// Create a TCP listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()
	
	// Register the greeting service
	pb.RegisterGreetingServiceServer(s, srv.NewServer())
	
	log.Printf("Server starting on port %s...", port)
	
	// Start serving
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
