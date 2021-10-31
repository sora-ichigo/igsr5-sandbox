package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	go func(ch chan int) {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 0)
			ch <- i
		}
		close(ch)
	}(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
