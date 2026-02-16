package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ElodinLaarz/grpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server implements the GreetingService
type server struct {
	pb.UnimplementedGreetingServiceServer
}

// SayHello implements the SayHello RPC method
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := req.GetName()
	log.Printf("Received request from client: name=%s", name)
	
	message := fmt.Sprintf("Hello, %s!", name)
	log.Printf("Sending response to client: message=%s", message)
	
	return &pb.HelloResponse{Message: message}, nil
}

func main() {
	// Create a TCP listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()
	
	// Register the greeting service
	pb.RegisterGreetingServiceServer(s, &server{})
	
	log.Printf("Server starting on port %s...", port)
	
	// Start serving
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
