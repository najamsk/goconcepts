package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	t := 0
	res := [][]int{}
	sort.Ints(nums)
	for k, v := range nums {
		if k > 0 && nums[k-1] == v {
			continue
		}
		l := k + 1
		r := len(nums) - 1
		for l < r {
			sum := v + nums[l] + nums[r]
			if sum == t {
				res = append(res, []int{v, nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}

				// return res
			} else if sum > t {
				r--
			} else {
				l++
			}

		}
	}

	return res
}

func main() {
	//https://leetcode.com/problems/3sum/
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
