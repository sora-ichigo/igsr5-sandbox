package main

import "fmt"

func main() {
	ans := fib(5)
	fmt.Printf("ans: %d\n", ans)
	fmt.Printf("ans: %d\n", fib1(50))

	memo := map[int]int{}
	fmt.Printf("ans: %d\n", fib2(50, memo))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func fib1(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func fib2(n int, memo map[int]int) int {
	if n < 2 {
		return n
	}
	if _, ok := memo[n]; !ok {
		memo[n] = fib2(n-1, memo) + fib2(n-2, memo)
	}

	return memo[n]
}
