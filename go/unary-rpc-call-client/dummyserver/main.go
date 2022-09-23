package main

import (
	"context"
	"fmt"
	"log"
	"net"

	example_pb "github.com/igsr5/unary-rpc-call-client/example"
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
	example_pb.RegisterTestServiceServer(s, newTestService())

	fmt.Printf("Starting server on port 8080\n")
	return s.Serve(lis)
}

func newTestService() example_pb.TestServiceServer {
	return &dummyserverImpl{}
}

type dummyserverImpl struct{}

func (d *dummyserverImpl) TestMethod(ctx context.Context, req *example_pb.ExampleRequest) (*example_pb.ExampleResponse, error) {
	return &example_pb.ExampleResponse{
		Name: "pong",
	}, nil
}
