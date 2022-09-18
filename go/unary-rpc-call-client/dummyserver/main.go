package main

import (
	"context"
	"fmt"
	"log"
	"net"

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
	RegisterTestServiceServer(s, newTestService())

	fmt.Printf("Starting server on port 8080\n")
	return s.Serve(lis)
}

func newTestService() TestServiceServer {
	return &dummyserverImpl{}
}

type dummyserverImpl struct{}

func (d *dummyserverImpl) TestMethod(ctx context.Context, req *ExampleRequest) (*ExampleResponse, error) {
	return &ExampleResponse{
		Name: "pong",
	}, nil
}
