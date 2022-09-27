package miniSubSum

import "math"

func MinimumSubSum(target int, items []int) int {
	sum := 0
	l := 0
	start := 0

	for end := 0; end < len(items); end++ {
		sum += items[end]
		for sum >= target {
			if l == 0 {
				l = end - start + 1
			}
			l = int(math.Min(float64(l), float64(end-start+1)))
			sum -= items[start]
			start++
		}
	}
	if l == 0 {
		return 0
	}
	return int(l)
}
