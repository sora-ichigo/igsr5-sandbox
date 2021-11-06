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

	ch = make(chan int, 1)
	ch <- 3

	// output: 3
	fmt.Printf("before close: %v\n", <-ch)

	ch <- 3
	close(ch)
	// output: 0
	fmt.Printf("after close: %v\n", <-ch)

	// output: panic: send on closed channel
	ch <- 4
}
