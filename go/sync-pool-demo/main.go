package syncpooldemo

import "sync"

func main() {
}

func arrayAppend() {
	var s []int

	for i := 0; i < 100; i++ {
		s = append(s, i)
	}
}

func arrayMake() {
	s := make([]int, 100)

	for i := 0; i < 100; i++ {
		s = append(s, i)
	}
}

func arrayPool() {
	pool := sync.Pool{
		New: func() interface{} {
			return make([]int, 100)
		},
	}

	for i := 0; i < 100; i++ {
		p := pool.Get().([]int)
		p = append(p, i)
		pool.Put(p)
	}
}
