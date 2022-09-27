package fixed

import "math"

// 1, 12, -5, -6, 50, 3
func MaxAverageSumSubset(items []int, k int) float64 {
	start := 0
	sum := 0
	max := math.Inf(-1)

	for end := 0; end < len(items); end++ {
		sum += items[end]

		//window size reached
		if end-start+1 == k {
			//calculate avg for window
			max = math.Max(max, float64(sum)/float64(k))
			//shift window
			sum -= items[start]
			start++
		}
	}
	return max
}
