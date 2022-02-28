package main

import "fmt"

func main() {
	var b struct {
		name string
		age  int
	}
	hoge1(&b)
	fmt.Println("b: ", b)

	hoge2(&b)
	fmt.Println("b: ", b)

	var c *struct {
		name string
		age  int
	}
	fmt.Printf("address(c)1: %p\n", c)
	hoge1(c)
	fmt.Println("c: ", c)
	fmt.Printf("address(c)3: %p\n", c)

	hoge2(c)
	fmt.Println("c: ", c)
	fmt.Printf("address(c)5: %p\n", c)
}

func hoge1(a *struct {
	name string
	age  int
}) {
	if a == nil {
		a = &struct {
			name string
			age  int
		}{}
	}

	fmt.Printf("address(c)2: %p\n", a)
	a.name = "hoge"
}

func hoge2(a *struct {
	name string
	age  int
}) {
	if a == nil {
		a = &struct {
			name string
			age  int
		}{}
	}

	fmt.Printf("address(c)4: %p\n", a)
	a.age = 3
}
