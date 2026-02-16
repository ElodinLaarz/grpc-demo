package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ElodinLaarz/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server
	log.Printf("Connecting to server at %s...", address)
	
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	
	log.Printf("Connected to server successfully")
	
	// Create a client
	c := pb.NewGreetingServiceClient(conn)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Make multiple calls to demonstrate the demo
	names := []string{"Alice", "Bob", "Charlie"}
	
	for _, name := range names {
		log.Printf("Requesting greeting for: %s", name)
		
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("Failed to call SayHello: %v", err)
		}
		
		log.Printf("Greeting: %s", r.GetMessage())
		
		// Small delay between requests
		time.Sleep(500 * time.Millisecond)
	}
	
	log.Printf("All requests completed successfully")
}
