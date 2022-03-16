package main

import (
	"fmt"

	userpb "github.com/igsr5/protosum/go/user"
)

func main() {
	fmt.Println("Hello")
	fmt.Printf("%#v", new(userpb.User))
}
