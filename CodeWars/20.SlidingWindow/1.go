package main

import (
	"fmt"
	dynamictree "sliding/dynamicTree"
	fixed "sliding/fixed"
	longestsubstring "sliding/longestSubstring"
	"sliding/miniSubSum"
)

func main() {
	fmt.Println("hello")
	res := fixed.MaxAverageSumSubset([]int{1, 12, -5, -6, 50, 3}, 4)
	fmt.Println(res)
	res = float64(dynamictree.Pick([]int{0, 1, 6, 6, 4, 4, 6}, 2))
	fmt.Println(res)
	l := longestsubstring.Longestsubstring("pwwkew")
	fmt.Println(l)
	m := miniSubSum.MinimumSubSum(7, []int{2, 3, 1, 2, 4, 3})

	// m := minisubsum.MinimumSubSum(7, []int{2, 3, 1, 2, 4, 3})
	fmt.Println("subSum=", m)
}
