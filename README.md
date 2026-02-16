# grpc-demo

A robust gRPC service demo in Go, demonstrating best practices for structure, validation, and testing.

This project implements a simple "Greeting Service" where a client sends a name and the server responds with a greeting. It goes beyond "Hello World" by including input validation and a full integration test suite.

## Architecture

- **Proto Definition**: `proto/greeting.proto`
    - Defines the `GreetingService` with a `SayHello` RPC.
- **Server**: `internal/server` & `cmd/server`
    - Implements the server logic with input validation.
    - Returns `InvalidArgument` if the name is empty.
- **Client**: `cmd/client`
    - Connects to the server and sends a sequence of requests.
- **Testing**: `internal/server/server_test.go`
    - Uses `bufconn` (in-memory connection) to run full end-to-end integration tests without network ports.

## Prerequisites

- Go 1.24 or later
- Protocol Buffers compiler (`protoc`)
- Make (optional, for ease of use)

## Setup

1.  Clone the repository:
    ```bash
    git clone https://github.com/ElodinLaarz/grpc-demo.git
    cd grpc-demo
    ```

2.  Install dependencies:
    ```bash
    go mod tidy
    ```

**Note:** All commands below should be run from the root of the project directory (`grpc-demo/`).

## Quick Start

The easiest way to run everything is using the `Makefile`.

### 1. Run Tests
Run the comprehensive test suite (Unit + Integration):
```bash
make test
```

### 2. Run Server & Client

**Terminal 1 (Server):**
Start the server. It will listen on `localhost:50051`.
```bash
make server
# OR
go run cmd/server/main.go
```

**Terminal 2 (Client):**
Run the client to send requests to the server.
```bash
make client
# OR
go run cmd/client/main.go
```

## Project Structure

```
.
├── cmd/
│   ├── client/    # Client entry point
│   └── server/    # Server entry point
├── internal/
│   └── server/    # Server implementation & Tests
├── proto/         # Protocol Buffer definitions
├── Makefile       # Build & Run automation
└── README.md
```

## Testing Strategy

This project uses **in-memory integration testing** for the gRPC service.

Instead of spinning up a server on a real TCP port (which can be flaky or conflict with other services), we use `google.golang.org/grpc/test/bufconn`. This creates an in-memory network listener that the client can dial directly.

This approach allows us to test the *entire* stack (interceptor chain, serialization, handler logic) safely and quickly.

See `internal/server/server_test.go` for the implementation.

## Regenerating Proto Files

If you modify `proto/greeting.proto`, regenerate the Go code from the project root:

```bash
make proto
# OR
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/greeting.proto
```
