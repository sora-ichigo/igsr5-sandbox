package main

import (
	"testing"
)

func BenchmarkCopySlice1(b *testing.B) {
	sliceA := setup()

	for i := 0; i < b.N; i++ {
		copySlice1(sliceA)
	}
}

func BenchmarkCopySlice2(b *testing.B) {
	sliceA := setup[string]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copySlice2(sliceA)
	}
}

func setup[T any]() []T {
	sliceA := []T{}
	for i := 0; i < 1000; i++ {
		sliceA = append(sliceA, T(any(i)))
	}

	return sliceA
}
