.PHONY: all proto server client test clean

# Default target
all: proto

# Generate proto files
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/greeting.proto

# Run server
server:
	go run cmd/server/main.go

# Run client
client:
	go run cmd/client/main.go

# Run tests
test:
	go test -v ./...

# Clean generated files
clean:
	rm -f proto/*.pb.go
