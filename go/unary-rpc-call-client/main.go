package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/igsr5/unary-rpc-call-client/example"
	"google.golang.org/grpc"
)

// service&rpc name, rpc request, host を受け取って rpc response を返す
func main() {
	ctx := context.Background()

	target := "localhost:8080"
	serviceName := "github.com.igsr5.example.TestService"
	methodName := "TestMethod"

	req, resp := getReqAndResp(ServiceName(serviceName), MethodName(methodName))

	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	fullMethodName := fmt.Sprintf("/%s/%s", serviceName, methodName)
	err = conn.Invoke(ctx, fullMethodName, req, resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp)
}

type ServiceName string
type MethodName string
type FullMethodName string

var ProtoMap = map[ServiceName]map[MethodName]struct {
	Req  proto.Message
	Resp proto.Message
}{
	"github.com.igsr5.example.TestService": {
		"TestMethod": {
			Req:  &example.ExampleRequest{},
			Resp: &example.ExampleResponse{},
		},
	},
}

func getReqAndResp(serviceName ServiceName, methodName MethodName) (proto.Message, proto.Message) {
	data, ok := ProtoMap[serviceName][methodName]
	if ok {
		return data.Req, data.Resp
	}

	return nil, nil
}
