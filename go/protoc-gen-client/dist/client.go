package main

import (
	"context"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type JSON []byte

func ServiceDescriptor(fds []*desc.FileDescriptor, serviceName string) (*desc.ServiceDescriptor, error) {
	for _, fd := range fds {
		sd := fd.FindService(serviceName)
		if sd != nil {
			return sd, nil
		}
	}

	return nil, errors.Errorf("failed to find service %s", serviceName)
}

func MethodDescriptor(sd *desc.ServiceDescriptor, methodName string) (*desc.MethodDescriptor, error) {
	md := sd.FindMethodByName(methodName)
	if md == nil {
		return nil, errors.Errorf("failed to find method %s", methodName)
	}

	return md, nil
}

func MethodFullName(sd *desc.ServiceDescriptor, md *desc.MethodDescriptor) string {
	return "/" + sd.GetFullyQualifiedName() + "/" + md.GetName()
}

type Client interface {
	Call(ctx context.Context, serviceName string, methodName string, in JSON, targetHost string) (JSON, error)
}

type clientImpl struct {
}

func NewClient() Client {
	return &clientImpl{}
}

func (c *clientImpl) Call(ctx context.Context, serviceFullName string, methodName string, in JSON, targetHost string) (JSON, error) {
	conn, err := grpc.DialContext(ctx, targetHost, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial")
	}

	fds, err := FileDescriptors()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get file descriptors")
	}

	sd, err := ServiceDescriptor(fds, serviceFullName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get service descriptor")
	}

	md, err := MethodDescriptor(sd, methodName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get method descriptor")
	}

	reqMd := md.GetInputType()
	respMd := md.GetOutputType()

	req := dynamic.NewMessageFactoryWithDefaults().NewDynamicMessage(reqMd)
	err = req.UnmarshalMergeJSON(in)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	resp := dynamic.NewMessageFactoryWithDefaults().NewDynamicMessage(respMd)

	err = conn.Invoke(ctx, MethodFullName(sd, md), req, resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to invoke")
	}

	out, err := resp.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal response")
	}

	return out, nil
}
