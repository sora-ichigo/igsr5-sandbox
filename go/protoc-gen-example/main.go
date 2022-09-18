package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pkg/errors"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("failed to run() %v", err)
	}
}

func run() error {
	req, err := parseReq(os.Stdin)
	if err != nil {
		return errors.Wrap(err, "failed to parse request")
	}

	resp, err := processReq(req)
	if err != nil {
		return errors.Wrap(err, "failed to process request")
	}

	err = emitResp(resp)
	if err != nil {
		return errors.Wrap(err, "failed to emitResp")
	}

	return nil
}

func parseReq(r io.Reader) (*plugin.CodeGeneratorRequest, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read request")
	}

	req := &plugin.CodeGeneratorRequest{}
	err = proto.Unmarshal(buf, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return req, nil
}

func processReq(req *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	files := make(map[string]*descriptor.FileDescriptorProto, len(req.ProtoFile))
	for _, f := range req.ProtoFile {
		files[f.GetName()] = f
	}

	var resp plugin.CodeGeneratorResponse
	for k, f := range files {
		out := strings.TrimPrefix(fmt.Sprintf("%s.dump", k), "src/")
		resp.File = append(resp.File, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(out),
			Content: proto.String(proto.MarshalTextString(f)),
		})
	}

	return &resp, nil
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
