package main

import (
	"fmt"
	"sort"
)

func boxes(a []int) int {
	sort.Ints(a)
	fmt.Println(a)
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return 0
	}

	steps := 0
	for i := 0; i < len(a); i++ {
		if (i+1) < len(a) && a[i] != a[i+1] {
			steps++
		}
	}

	return steps + 1
}

func main() {
	fmt.Println(boxes([]int{1}))
	fmt.Println(boxes([]int{2, 3, 4, 3, 3}))
	fmt.Println(boxes([]int{}))
}
