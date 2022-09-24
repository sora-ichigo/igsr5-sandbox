package main

import (
	"context"
	"fmt"
	"log"
	"net"

	example_pb "github.com/igsr5/protoc-gen-client/example"
	grpc "google.golang.org/grpc"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	// Create a new server.
	s := grpc.NewServer()
	example_pb.RegisterExampleServiceServer(s, newExampleService())

	fmt.Printf("Starting server on port 8080\n")
	return s.Serve(lis)
}

func newExampleService() example_pb.ExampleServiceServer {
	return &dummyserverImpl{}
}

type dummyserverImpl struct{}

func (d *dummyserverImpl) Echo(ctx context.Context, req *example_pb.Example) (*example_pb.Example, error) {
	return &example_pb.Example{
		Name: "pong",
	}, nil
}
