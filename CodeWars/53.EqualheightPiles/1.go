package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/discuss/interview-question/364618/
// Alexa is given n piles of equal or unequal heights.
// In one step, Alexa can remove any number of boxes from the pile which has the maximum height and try to make it equal to the one which is just lower than the maximum height of the stack. Determine the minimum number of steps required to make all of the piles equal in height.
// Input: piles = [5, 2, 1]
// Output: 3
// Explanation:
// Step 1: reducing 5 -> 2 [2, 2, 1]
// Step 2: reducing 2 -> 1 [2, 1, 1]
// Step 3: reducing 2 -> 1 [1, 1, 1]
// So final number of steps required is 3.

func makeEqual(boxes []int) int {
	sort.Ints(boxes)
	steps := 0

	for i := 0; i < len(boxes); i++ {
		if (i+1) < len(boxes) && boxes[i+1] != boxes[i] {
			steps += len(boxes) - i - 1
		}
	}
	return steps

}
func main() {
	// fmt.Println(makeEqual([]int{2, 3, 4, 3, 3}))
	// fmt.Println(makeEqual([]int{43, 247}))
	fmt.Println(makeEqual([]int{4, 3, 2, 1}))
	fmt.Println(makeEqual([]int{3, 2, 1}))
	fmt.Println(makeEqual([]int{5, 2}))
	fmt.Println(makeEqual([]int{5}))
	// fmt.Println(makeEqual([]int{150, 210, 210, 80, 110}))
}
