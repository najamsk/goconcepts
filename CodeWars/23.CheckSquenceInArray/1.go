package main

import "fmt"

func test(nums, squence []int) bool {
	i := 0
	for _, v := range nums {
		if v == squence[i] {
			i++
		}
		if i == len(squence) {
			return true
		}
	}
	return false
}

func main() {
	v := test([]int{1, 2, 3, 4}, []int{1, 3})
	fmt.Println(v)
}
