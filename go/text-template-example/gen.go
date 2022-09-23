package main

import (
	"encoding/json"
	"fmt"
)

type PersonV2 struct {
	Name NameV2
}

type NameV2 struct {
	First string
	Last  string
}

func main() {
	v := []byte{}

	v = append(v, 123)

	v = append(v, 34)

	v = append(v, 78)

	v = append(v, 97)

	v = append(v, 109)

	v = append(v, 101)

	v = append(v, 34)

	v = append(v, 58)

	v = append(v, 123)

	v = append(v, 34)

	v = append(v, 70)

	v = append(v, 105)

	v = append(v, 114)

	v = append(v, 115)

	v = append(v, 116)

	v = append(v, 34)

	v = append(v, 58)

	v = append(v, 34)

	v = append(v, 74)

	v = append(v, 111)

	v = append(v, 104)

	v = append(v, 110)

	v = append(v, 34)

	v = append(v, 44)

	v = append(v, 34)

	v = append(v, 76)

	v = append(v, 97)

	v = append(v, 115)

	v = append(v, 116)

	v = append(v, 34)

	v = append(v, 58)

	v = append(v, 34)

	v = append(v, 68)

	v = append(v, 111)

	v = append(v, 101)

	v = append(v, 34)

	v = append(v, 125)

	v = append(v, 125)

	fmt.Println("Hello, World!")
	var p PersonV2
	json.Unmarshal(v, &p)
	fmt.Printf("v = %v", p)
}
