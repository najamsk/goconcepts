package main

import (
	"fmt"
)

func simpleSort(a []int) []int {
	i := 0
	for i < len(a) {
		if a[i] != i+1 {
			t := a[i] - 1
			a[t], a[i] = a[i], a[t]
		} else {
			i++
		}
	}
	return a
}

func missingNumber(a []int) int {
	i := 0
	for i < len(a) {
		if a[i] != i+1 {
			return i + 1
		} else {
			i++
		}
	}
	return len(a)
}

func missingPositiveNumber(a []int) int {
	i := 0
	for i < len(a) {
		if a[i] != i && a[i] < len(a) && !(a[i] < 0) {
			t := a[i]
			a[t], a[i] = a[i], a[t]
		} else {
			i++
		}
	}

	fmt.Println(a)
	for k := range a {
		if a[k] != k {
			return k
		}
	}

	//fixed
	return 0

}

func main() {
	fmt.Println(simpleSort([]int{3, 1, 5, 4, 2}))
	fmt.Println(missingNumber([]int{1, 2, 3, 5}))
	fmt.Println(missingPositiveNumber([]int{4, 0, 3, 1}))
	fmt.Println(missingPositiveNumber([]int{3, 0, -2, 1, 2}))
}
