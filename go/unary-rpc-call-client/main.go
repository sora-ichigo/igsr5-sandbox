package main

import (
	"context"
	"fmt"
	"log"

	"github.com/igsr5/unary-rpc-call-client/example"
	"google.golang.org/grpc"
)

// service&rpc name, rpc request, host を受け取って rpc response を返す
func main() {
	ctx := context.Background()

	req := example.ExampleRequest{}
	target := "localhost:8080"
	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	methodName := "/github.com.igsr5.example.TestService/TestMethod"

	var resp example.ExampleResponse
	err = conn.Invoke(ctx, methodName, &req, &resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp.GetName())
}
