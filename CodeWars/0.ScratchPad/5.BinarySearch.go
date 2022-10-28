package main

import "fmt"

func search(nums []int, t int) int {
	r := len(nums) - 1
	l := 0

	for l < r {
		mid := (l + r) / 2
		if nums[mid] == t {
			return mid
		} else if nums[mid] < t {
			l = mid + 1
		} else {
			r = mid
		}
	}

	fmt.Printf("l:%v, r:%v \n", l, r)
	if l <= r && nums[l] == t {
		return l
	}
	return -1
}

func main() {
	fmt.Println(search([]int{1, 3, 4, 8, 20}, 5))
}
