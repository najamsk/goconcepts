package main

import "fmt"

func calStep2(n int64) int64 {
	steps := make([]int64, n)
	steps[0] = 1
	steps[1] = 2
	steps[2] = 4
	idx := int64(3)
	for i := idx; i < n && n > 3; i++ {
		steps[i] = steps[i-1] + steps[i-2] + steps[i-3]
	}
	return steps[len(steps)-1]
}

func main() {
	fmt.Println(calStep2(3))
}
