package main

import (
	"fmt"
)

type planet struct {
	x int
	r int
}

func findGaps(planets []planet) [][]int {
	ranges := [][]int{}
	lines := make([]int, 20)

	for _, v := range planets {
		left := v.x - (v.r - 1)
		right := v.x + (v.r - 1)
		for i := left; i <= right; i++ {
			lines[i] = 1
		}
	}

	fmt.Println(lines)
	right := 0
	for right < len(lines) {

		if lines[right] == 0 {
			end := right
			for (end+1) < len(lines) && lines[end+1] == 0 {
				end++
			}
			cr := []int{right, end}
			ranges = append(ranges, cr)
			right = end + 1
		} else {
			right++
		}
	}
	return ranges

}

func main() {
	fmt.Println(findGaps([]planet{{0, 1}, {x: 3, r: 2}, {x: 5, r: 1}}))
}
