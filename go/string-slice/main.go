package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"

	arr := []string{}

	arr = strings.Split(str, ",")

	for _, a := range arr {
		f, _ := regexp.MatchString("(?)image/webp|image/apng", a)
		fmt.Println(f)
	}

	fmt.Println(arr)
}
