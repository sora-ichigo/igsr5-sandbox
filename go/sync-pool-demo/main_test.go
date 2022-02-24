package syncpooldemo

import (
	"sync"
	"testing"
)

// スライスを append のたびに拡張
func BenchmarkArrayAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s []int

		for j := 0; j < 100; j++ {
			s = append(s, j)
		}

	}
}

// 固定長のスライドを初めに確保
func BenchmarkArrayMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 100)

		for j := 0; j < 100; j++ {
			s = append(s, j)
		}

	}
}

var count = 100000

// 毎回のループで make([]int, count) を呼び出す
func BenchmarkWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < count; j++ {
			_ = make([]int, count)
		}
	}
}

// ループ前に pool を確保し、ループ内で使い回す
func BenchmarkWithPool(b *testing.B) {
	pool := sync.Pool{
		New: func() interface{} {
			return make([]int, count)
		},
	}

	for i := 0; i < b.N; i++ {
		for j := 0; j < count; j++ {
			p := pool.Get().([]int)
			pool.Put(p)
		}
	}
}
