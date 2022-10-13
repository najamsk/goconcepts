package main

import (
	"fmt"
)

func maxi(a []int) int {
	max := 0
	mN := make(map[int]int, 0)

	for _, v := range a {
		if v < 0 {
			mN[v]++
		}
	}
	for _, v := range a {
		if (v > max) && (v > 0) {
			neg := v * -1
			if _, ok := mN[neg]; ok {
				max = v
			}
		}
	}
	return max
}

func main() {
	fmt.Println(maxi([]int{5, 1, 20, -10, -5, 2, 10}))
}
