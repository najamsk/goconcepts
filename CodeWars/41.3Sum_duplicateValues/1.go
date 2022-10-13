package main

import (
	"fmt"
	"sort"
)

func threeSum(a []int, t int) []int {
	res := []int{}
	//sort input array
	sort.Ints(a)
	fmt.Println(a)
	for k, v := range a {
		if k > 0 && a[k-1] == v {
			continue
		}
		// fmt.Printf("%v", v)
		l := k + 1
		r := len(a) - 1
		for l < r {
			if a[l] == v {
				l++
			}
			sum := v + a[l] + a[r]
			if sum == t {
				res = append(res, v, a[l], a[r])
				return res
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
	fmt.Println(threeSum([]int{-3, 3, 4, -3, 1, 2}, 0))
}
