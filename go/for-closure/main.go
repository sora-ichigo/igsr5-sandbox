package main

import (
	"fmt"
	"time"
)

func LoopWithClosure() {
	f := make([]func(), 3)

	for i := 0; i < 3; i++ {
		// このループ内で値が参照されないと、スコープ外にでてしまい、
		// Go がいい感じに最後の値を利用しちゃう
		f[i] = func() {
			fmt.Println(i)
		}
	}

	fmt.Println("closure")
	for _, f := range f {
		f()
	}
}

func LoopWithMethod() {
	f := make([]func(), 3)

	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	fmt.Println("method")
	for _, f := range f {
		f()
	}
}

func LoopWithNewVal() {
	f := make([]func(), 3)

	for i := 0; i < 3; i++ {
		i := i
		f[i] = func() {
			fmt.Println(i)
		}
	}

	fmt.Println("method")
	for _, f := range f {
		f()
	}
}

func main() {
	for i, v := range []int{1, 2, 3} {
		t := v
		fmt.Printf("i: %d, v: %v\n", i, &v)
		go func() {
			fmt.Println(t)
		}()
	}

	time.Sleep(1 * time.Second)
}
