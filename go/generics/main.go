package main

import "fmt"

type Number interface {
	int64 | float64
}

type Set[V comparable] struct {
	m map[V]struct{}
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}

	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}

	return s
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	ints := map[string]int64{"a": 1, "b": 2, "c": 3}
	floats := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}

	fmt.Printf("Non-Generic Sums: %v and %v\n", SumInts(ints), SumFloats(floats))
	fmt.Printf("Generic Sums: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))

	sets := Set[int]{map[int]struct{}{1: {}, 2: {}, 3: {}}}
	fmt.Printf("Set: %v\n", sets)
	sets1 := Set[string]{map[string]struct{}{"1": {}, "2": {}, "3": {}}}
	fmt.Printf("Set: %v\n", sets1)
}
