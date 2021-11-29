package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	url, err := url.Parse("https://example.com/image.png")
	//url, err := url.Parse("")
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	fmt.Println(url)
	path := url.Path
	fmt.Println(path)
	fmt.Println(strings.Split(path, "/")[len(strings.Split(path, "/"))-1])
}
