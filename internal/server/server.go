package server

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ElodinLaarz/grpc-demo/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server implements the GreetingService
type Server struct {
	pb.UnimplementedGreetingServiceServer
}

// NewServer creates a new GreetingService server
func NewServer() *Server {
	return &Server{}
}

// SayHello implements the SayHello RPC method
func (s *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := req.GetName()
	log.Printf("Received request from client: name=%s", name)

	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "name cannot be empty")
	}
	
	message := fmt.Sprintf("Hello, %s!", name)
	log.Printf("Sending response to client: message=%s", message)
	
	return &pb.HelloResponse{Message: message}, nil
}
