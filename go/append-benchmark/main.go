package main

import (
	"fmt"
	"strconv"
)

func main() {
	sliceA := []string{}
	for i := 0; i < 100; i++ {
		sliceA = append(sliceA, strconv.FormatInt(int64(i), 10))
	}

	sliceB := copySlice2(sliceA)
	fmt.Printf("%#v", sliceB)
}

func copySlice1(a []string) []string {
	b := []string{}
	for _, v := range a {
		b = append(b, v)
	}

	return b
}

func copySlice2(a []string) []string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = v
	}

	return b
}
