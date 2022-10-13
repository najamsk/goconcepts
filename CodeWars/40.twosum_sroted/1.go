package main

import (
	"fmt"
)

func twosum(a []int, t int) []int {
	res := []int{}
	l := 0
	r := len(a) - 1

	for l < r {
		sum := a[l] + a[r]
		if sum == t {
			res = append(res, a[l], a[r])
			return res
		} else if sum > t {
			r--
		} else {
			l++
		}
	}
	return res
}

func main() {
	fmt.Println(twosum([]int{1, 3, 4, 5, 7, 10, 11}, 9))
}
