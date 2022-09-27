package longestsubstring

import (
	"math"
)

func Longestsubstring(s string) int {
	//pwwkew
	l := 0
	start := 0
	m := make(map[byte]bool)

	for end := range s {
		c := s[end]
		if !m[c] {
			m[c] = true
			l = int(math.Max(float64(l), float64(end-start+1)))
		} else {
			for m[c] {
				m[s[start]] = false
				start++
			}
			m[c] = true
		}
	}
	return int(l)
}
