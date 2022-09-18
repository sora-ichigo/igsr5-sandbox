package main

import (
	"log"
	"os"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("failed to run() %v", err)
	}
}

func run() error {
	resp := &plugin.CodeGeneratorResponse{
		File: []*plugin.CodeGeneratorResponse_File{
			{
				Name:    proto.String("foo.txt"),
				Content: proto.String("hello world"),
			},
		},
	}
	err := emitResp(resp)
	if err != nil {
		return errors.Wrap(err, "failed to emitResp")
	}

	return nil
}

func emitResp(resp *plugin.CodeGeneratorResponse) error {
	buf, err := proto.Marshal(resp)
	if err != nil {
		return errors.Wrap(err, "failed to marshal response")
	}

	_, err = os.Stdout.Write(buf)
	if err != nil {
		return errors.Wrap(err, "failed to write response")
	}

	return nil
}
