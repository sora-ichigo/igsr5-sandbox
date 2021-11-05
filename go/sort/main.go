package main

import "fmt"

func main() {
	arr := []int{1, 4, 5, 6, 2}
	fmt.Println("bubbleSort")
	fmt.Printf("before: %v\n", arr)
	fmt.Printf("after: %v\n", bubbleSort(arr))

	fmt.Println("")

	arr = []int{1, 4, 5, 6, 2, 11, 55, 3232, 5}
	fmt.Println("selectionSort")
	fmt.Printf("before: %v\n", arr)
	fmt.Printf("after: %v\n", selectionSort(arr))

	fmt.Println("")

	arr = []int{1, 4, 5, 6, 2, 11, 55, 3232, 5}
	fmt.Println("insertionSort")
	fmt.Printf("before: %v\n", arr)
	fmt.Printf("after: %v\n", insertionSort(arr))

	fmt.Println("")

	arr = []int{1, 4, 5, 6, 2, 11, 55, 3232, 5}
	fmt.Println("quickSort")
	fmt.Printf("before: %v\n", arr)
	fmt.Printf("after: %v\n", quickSort(arr))

	fmt.Println("")

	arr = []int{1, 4, 5, 6, 2, 11, 55, 3232, 5}
	fmt.Println("mergeSort")
	fmt.Printf("before: %v\n", arr)
	fmt.Printf("after: %v\n", mergeSort(arr))

	fmt.Println("")

	arr = []int{1, 4, 5, 6, 2, 11, 55, 3232, 5}
	fmt.Println("shellSort")
	fmt.Printf("before: %v\n", arr)
	fmt.Printf("after: %v\n", shellSort(arr))

	fmt.Println("")

	arr = []int{1, 4, 5, 6, 2, 11, 55, 3232, 5}
	fmt.Println("heapSort")
	fmt.Printf("before: %v\n", arr)
	fmt.Printf("buildHeap: %v\n", buildHeap(arr))
	fmt.Printf("after: %v\n", heapSort(arr))
}

func bubbleSort(arr []int) []int {
	for j, _ := range arr {
		for i, _ := range arr[:(len(arr) - j - 1)] {
			if arr[i] > arr[i+1] {
				tmp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
			}
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	var p, min int

	for j, _ := range arr {
		p = j
		min = arr[p]
		for i, n := range arr {
			if i > j && min > n {
				p = i
			}
		}
		tmp := arr[j]
		arr[j] = arr[p]
		arr[p] = tmp
	}
	return arr
}

func insertionSort(arr []int) []int {
	var key int

	for j := 1; j < len(arr); j++ {
		key = arr[j]
		i := j - 1
		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i--
		}
		arr[i+1] = key
	}
	return arr
}

func quickSort(arr []int) []int {
	var left, right []int
	if len(arr) > 1 {
		x := arr[0]
		for _, a := range arr[1:] {
			if a < x {
				left = append(left, a)
			} else {
				right = append(right, a)
			}
		}
		left = quickSort(left)
		right = quickSort(right)
		arr = append(append(left, x), right...)
	}

	return arr
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	first := mergeSort(arr[len(arr)/2:])
	second := mergeSort(arr[:len(arr)/2])

	return merge(first, second)
}

func merge(a, b []int) []int {
	final := []int{}

	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for i < len(a) {
		final = append(final, a[i])
		i++
	}

	for j < len(b) {
		final = append(final, b[j])
		j++
	}

	return final
}

func shellSort(arr []int) []int {
	h := 4

	for h > 1 {
		for i := 0; i < h; i++ {
			b := make([]int, len(arr)/h+1)
			cnt := 0

			for j := 0; j < len(arr)/h+1; j++ {
				if i+h*j < len(arr) {
					b[j] = arr[i+h*j]
					cnt++
				}
			}
			c := insertionSort(b[:cnt])

			for j := 0; j < cnt; j++ {
				if i+h*j < len(arr) {
					arr[i+h*j] = c[j]
				}
			}
		}
		h /= 2
	}

	arr = insertionSort(arr)
	return arr
}

func buildHeap(arr []int) []int {
	last := len(arr) - 1
	parent := last

	for parent >= 0 {
		children := [2]int{parent * 2, parent*2 + 1}

		// no children
		if children[0] > last {
			parent--
			continue
		}

		// only children[0]
		if children[1] > last {
			if arr[children[0]] > arr[parent] {
				tmp := arr[children[0]]
				arr[children[0]] = arr[parent]
				arr[parent] = tmp

				parent = children[0]
			}
			parent--
			continue
		}

		// both exists children 0, 1
		if arr[children[0]] > arr[children[1]] {
			if arr[children[0]] > arr[parent] {
				tmp := arr[children[0]]
				arr[children[0]] = arr[parent]
				arr[parent] = tmp

				parent = children[0]
				continue
			}
		} else {
			if arr[children[1]] > arr[parent] {
				tmp := arr[children[1]]
				arr[children[1]] = arr[parent]
				arr[parent] = tmp

				parent = children[1]
				continue
			}
		}
		parent--
	}

	return arr
}

func heapSort(arr []int) []int {
	resp := []int{}
	last := len(arr)

	for ; last > 0; last-- {
		arr := buildHeap(arr[:last])
		resp = append([]int{arr[0]}, resp...)
		fmt.Println("resp:", resp)
		arr[0] = arr[last-1]
	}

	return resp
}
