package dynamictree

import "math"

func Pick(trees []int, k int) int {
	m := make(map[int]bool)
	start := 0
	max := 0
	for end, tree := range trees {
		if len(m) < k && !m[tree] {
			m[tree] = true
			max = int(math.Max(float64(max), float64(end-start+1)))
		} else if m[tree] {
			// map is full and tree is already in map
			max = int(math.Max(float64(max), float64(end-start+1)))
		} else {
			m = make(map[int]bool)
			m[trees[end-1]] = true
			m[tree] = true
			start = end - 1

			for trees[start] == trees[start-1] {
				start--
			}

			max = int(math.Max(float64(max), float64(end-start+1)))
		}
	}
	return max

}
