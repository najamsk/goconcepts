package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	l := 0
	r := 1
	p := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			pro := prices[r] - prices[l]
			p = int(math.Max(float64(p), float64(pro)))
		} else {
			l = r
		}
		r++
	}
	return p
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
