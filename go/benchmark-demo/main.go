package main

import (
	"fmt"
	"strconv"
)

func ItoaByFmt(n int) {
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, fmt.Sprint(i))
	}
}

func ItoaByStrconv1(n int) {
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, strconv.Itoa(i))
	}
}

func ItoaByStrconv2(n int) {
	s := make([]string, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, strconv.Itoa(i))
	}
}
