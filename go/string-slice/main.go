package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	Type string
	Q    float64
}

var str = "application/xhtml+xml,application/xml;q=0.9,image/avif,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"

func main() {
	arr := strings.Split(str, ",")
	data := make([]Data, len(arr))

	for i, aaa := range arr {
		aa := strings.Split(aaa, ";")

		mimeType := strings.Trim(aa[0], " ")
		q := getQFromAcceptString(aa)
		data[i] = Data{Type: mimeType, Q: q}
	}

	sort.SliceStable(data, func(i, j int) bool { return data[i].Q > data[j].Q })

	fmt.Printf("%#v", data)
}

func getQFromAcceptString(arr []string) float64 {
	for _, a := range arr {
		if strings.HasPrefix(a, "q=") {
			q, err := strconv.ParseFloat(a[2:], 64)
			if err != nil {
				return 1.0
			}

			return q
		}
	}

	return 1.0
}
