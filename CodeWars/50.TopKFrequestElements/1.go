package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	res := []int{}
	m := make(map[int]int, 0)

	//count occurances
	for _, v := range nums {
		m[v]++
	}
	fmt.Println(m)
	// fmt.Println("len of nums", len(nums))
	//bucket arr
	b := make([][]int, len(nums)+1)

	// fmt.Println(b)
	for k, v := range m {
		b[v] = append(b[v], k)
	}
	fmt.Println(b)

	for i := len(b) - 1; i >= 0 && k > 0; i-- {
		if len(b[i]) > 0 {
			res = append(res, b[i]...)
			k = k - len(b[i])
		}
	}

	// fmt.Println(b)

	return res
}
func main() {
	/* fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1)) */
	fmt.Println(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 1))

}
