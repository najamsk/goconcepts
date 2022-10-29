package main

import "fmt"

func test(a []int, t int) int {
	l := 0
	r := len(a) - 1

	for l <= r {
		mid := (l + r) / 2
		// fmt.Println("mid:", mid)
		if a[mid] == t {
			return mid
		} else if a[mid] < t {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}
	return -1
}

func main() {
	fmt.Println(test([]int{2, 3, 6, 7, 8, 14, 22}, 30))
	fmt.Println(test([]int{2, 3, 6, 7, 8, 14, 22}, -1))
	fmt.Println(test([]int{2, 3, 6, 7, 8, 14, 22}, 22))
	fmt.Println(test([]int{2, 3, 6, 7, 8, 14, 22}, 2))
	fmt.Println(test([]int{2, 3, 6, 7, 8, 14, 22}, 3))
	fmt.Println(test([]int{2, 3, 6, 7, 8, 14, 22}, 14))

	fmt.Println(test([]int{2, 3, 6, 7, 8, 14, 22, 30}, 5))
}
