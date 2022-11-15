package main

import (
	"fmt"
)

func test(a, b []int) []int {
	res := []int{}

	for len(a) > 0 && len(b) > 0 {
		if a[0] < b[0] {
			res = append(res, a[0])
			a = a[1:]
		} else {
			res = append(res, b[0])
			b = b[1:]
		}
	}

	if len(a) > 0 {
		res = append(res, a...)
	}
	if len(b) > 0 {
		res = append(res, b...)
	}

	return res

}

func main() {
	fmt.Println(test([]int{1, 3, 5, 7}, []int{2, 4, 6, 8, 9, 10}))
}
