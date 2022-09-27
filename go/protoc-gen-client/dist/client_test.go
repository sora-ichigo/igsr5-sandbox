package main

import (
	"context"
	"testing"
)

func TestServiceDescriptor(t *testing.T) {
	t.Parallel()
	fds, err := FileDescriptors()
	if err != nil {
		t.Fatalf("failed to get file descriptors: %v", err)
	}

	sd, err := ServiceDescriptor(fds, "example.ExampleService")
	if err != nil {
		t.Fatal(err)
	}

	if sd == nil {
		t.Fatalf("failed to find service")
	}
}

func TestMethodDescriptor(t *testing.T) {
	t.Parallel()
	fds, err := FileDescriptors()
	if err != nil {
		t.Fatalf("failed to get file descriptors: %v", err)
	}

	sd, err := ServiceDescriptor(fds, "example.ExampleService")
	if err != nil {
		t.Fatal(err)
	}

	if sd == nil {
		t.Fatalf("failed to find service")
	}

	md, err := MethodDescriptor(sd, "Echo")
	if err != nil {
		t.Fatal(err)
	}

	if md == nil {
		t.Fatalf("failed to find method")
	}
}

func TestCall(t *testing.T) {
	t.Parallel()
	c := NewClient()

	_, err := c.Call(context.Background(), "example.ExampleService", "Echo", []byte(`{"name": "aa", "id": 1}`), "localhost:8080")
	if err != nil {
		t.Fatal(err)
	}

}
