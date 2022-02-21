package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"

	arr := []string{}

	arr = strings.Split(str, ",")

	fmt.Println(arr)
}
