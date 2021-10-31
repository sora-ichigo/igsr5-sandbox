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