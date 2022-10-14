package main

import "fmt"

func lift(arr []int) int {
	steps := 0

	if len(arr) < 2 {
		return steps
	}
	l := len(arr)
	for k := 1; k < l; k++ {
		if arr[l-k-1] > arr[l-k] {
			steps = steps + k
		}

	}
	return steps
}

func main() {
	fmt.Println(lift([]int{6, 5, 2, 1}))
}
