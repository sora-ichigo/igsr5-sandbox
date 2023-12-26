package main

import "fmt"

func main() {
	hoge := "hoge"
	fuga := "fuga"
	hello(hoge, fuga)
	hello(fuga, hoge) // build が通ってしまう 🙅

	hogev2 := Hoge("hoge")
	fugav2 := Fuga("fuga")
	hellov2(hogev2, fugav2)
	hellov2(fugav2, hogev2) // build が通らない 🙆

}

func hello(hoge, fuga string) {
	fmt.Println(hoge, fuga)
}

type Hoge string
type Fuga string

func hellov2(hoge Hoge, fuga Fuga) {
	fmt.Println(hoge, fuga)
}
