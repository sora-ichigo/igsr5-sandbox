package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/tree"
	"golang.org/x/tour/wc"
)

/*
A Tour of Go

Ref: https://qiita.com/rock619/items/f412d1f870a022c142d0#exercise-loops-and-functions
*/

func main() {
	// sqrt
	fmt.Println(sqrt(2))
	br()

	// pic
	pic.Show(getPic)
	br()

	// WordCount
	wc.Test(wordCount)
	br()

	// fibonacchi
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	br()

	// IPAddr
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip.String())
	}
	br()

	// sqrt error
	fmt.Println(sqrt(-2))
	br()

	// tree
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
}

func br() {
	fmt.Println("")
	fmt.Println("===========================================")
	fmt.Println("")
}

type ErrEegativeSqrt float64

func (e ErrEegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrEegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func getPic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8((x + y) / 2)
		}
	}
	return pic
}

func wordCount(s string) map[string]int {
	words := strings.Fields(s)

	resp := map[string]int{}

	for _, word := range words {
		resp[word]++
	}

	return resp
}

func fib() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a+b
		return a
	}
}

// IPAddr is ip address. e.g.  { 192, 168, 0, 1 }
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}
