package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var b struct {
		Url string
	}
	str := `{"url":"hoge"}`
	if err := json.Unmarshal([]byte(str), &b); err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("%#v\n", b)
}
