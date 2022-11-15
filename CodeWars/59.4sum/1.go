package main

import (
	"fmt"
	"sort"
)

func foursum(nums []int, t int) []int {
	res := []int{}
	quad := []int{}
	sort.Ints(nums)
	fmt.Println(nums)

	var ksum func(k, start, target int)
	ksum = func(k, start, target int) {
		if k != 2 {
			for i := start; i < len(nums)-k; i++ {
				if i > start && nums[i] == nums[i-1] {
					continue
				}
				quad = append(quad, nums[i])
				ksum(k-1, i+1, target-nums[i])
				//pop
				quad = quad[1:]
				return
			}

		}
		//base case two sum
		l, r := start, len(nums)-1
		for l < r {
			if nums[l]+nums[r] < target {
				l = l + 1
			} else if nums[l]+nums[r] > target {
				r = r - 1
			} else {
				res = append(res, nums[l], nums[r])
				res = append(res, quad...)
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}

	}
	ksum(4, 0, t)

	return res
}

func main() {
	fmt.Println(foursum([]int{1, 0, -1, 0, -2, 2}, 0))
}
