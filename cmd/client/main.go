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

	// Make multiple calls to demonstrate the demo
	names := []string{"Alice", "Bob", "Charlie", "Dave", "Eve"}
	
	// Loop indefinitely
	for {
		for _, name := range names {
			// Create a context with timeout for this request
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			
			log.Printf("Requesting greeting for: %s", name)
			
			r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
			if err != nil {
				log.Printf("Failed to call SayHello: %v", err)
			} else {
				log.Printf("Greeting: %s", r.GetMessage())
			}
			
			cancel()
			
			// Sleep 3 seconds between requests
			time.Sleep(3 * time.Second)
		}
	}
}
