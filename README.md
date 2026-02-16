# grpc-demo

Simple demo of creating a gRPC service in Golang. This demo demonstrates a basic client-server architecture where a client sends greeting requests to a server, and the server responds with greeting messages. Logs are displayed on both sides to show the communication flow.

## Architecture

- **Proto Definition**: `proto/greeting.proto` - Defines the gRPC service with a `SayHello` RPC method
- **Server**: `server.go` - Implements the gRPC server that listens on port 50051
- **Client**: `client.go` - Implements the gRPC client that sends requests to the server

## Prerequisites

- Go 1.24 or later
- Protocol Buffers compiler (protoc)

## Setup

1. Clone the repository:
```bash
git clone https://github.com/ElodinLaarz/grpc-demo.git
cd grpc-demo
```

2. Install dependencies:
```bash
go mod download
```

## Running the Demo

### Terminal 1: Start the Server

```bash
go run server.go
```

You should see output like:
```
2024/02/16 12:00:00 Server starting on port :50051...
```

### Terminal 2: Run the Client

In a separate terminal, run:

```bash
go run client.go
```

You should see output like:
```
2024/02/16 12:00:01 Connecting to server at localhost:50051...
2024/02/16 12:00:01 Connected to server successfully
2024/02/16 12:00:01 Sending request to server: name=Alice
2024/02/16 12:00:01 Received response from server: message=Hello, Alice!
2024/02/16 12:00:01 Sending request to server: name=Bob
2024/02/16 12:00:01 Received response from server: message=Hello, Bob!
2024/02/16 12:00:01 Sending request to server: name=Charlie
2024/02/16 12:00:01 Received response from server: message=Hello, Charlie!
2024/02/16 12:00:02 All requests completed successfully
```

The server terminal will show corresponding logs:
```
2024/02/16 12:00:01 Received request from client: name=Alice
2024/02/16 12:00:01 Sending response to client: message=Hello, Alice!
2024/02/16 12:00:01 Received request from client: name=Bob
2024/02/16 12:00:01 Sending response to client: message=Hello, Bob!
2024/02/16 12:00:01 Received request from client: name=Charlie
2024/02/16 12:00:01 Sending response to client: message=Hello, Charlie!
```

## Regenerating Proto Files

If you modify the proto definition, regenerate the Go code:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/greeting.proto
```

## What's Happening

1. The **server** starts and listens on port 50051
2. The **client** connects to the server
3. The client sends three greeting requests (for Alice, Bob, and Charlie)
4. The server receives each request, logs it, creates a response, and sends it back
5. The client receives each response and logs it
6. Both sides show detailed logs of the message exchange

This demonstrates the basic gRPC request-response pattern with full visibility into the communication flow.
