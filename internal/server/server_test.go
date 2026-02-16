package server

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/ElodinLaarz/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, NewServer())
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestSayHello(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.NewClient("passthrough://bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetingServiceClient(conn)

	tests := []struct {
		name    string
		reqName string
		wantMsg string
		wantErr codes.Code
	}{
		{
			name:    "Valid Name",
			reqName: "Gemini",
			wantMsg: "Hello, Gemini!",
			wantErr: codes.OK,
		},
		{
			name:    "Empty Name",
			reqName: "",
			wantMsg: "",
			wantErr: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: tt.reqName})
			
			// Check error code
			if status.Code(err) != tt.wantErr {
				t.Errorf("SayHello() error code = %v, want %v", status.Code(err), tt.wantErr)
			}

			// If we expected an error, we don't check the message
			if tt.wantErr != codes.OK {
				return
			}

			if resp.GetMessage() != tt.wantMsg {
				t.Errorf("SayHello() message = %v, want %v", resp.GetMessage(), tt.wantMsg)
			}
		})
	}
}
