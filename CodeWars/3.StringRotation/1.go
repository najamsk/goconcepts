package main

import (
	"fmt"
	"math"
)

/* Input : s = "GeeksforGeeks"
       d = 2
Output : Left Rotation  : "eksforGeeksGe"
        Right Rotation : "ksGeeksforGee"


Input : s = "qwertyu"
       d = 2
Output : Left rotation : "ertyuqw"
        Right rotation : "yuqwert" */

// func StringRightRotation(val string, d int) string {
// 	if d < 0 || len(val) == 0 {
// 		return val
// 	}
// 	// break string into runes
// 	r := len(val) - d%len(val)
// 	left := val[r:]
// 	right := val[:r]
// 	val = left + right
// 	return val
// }
// func StringLeftRotation(val string, d int) string {
// 	if d < 0 || len(val) == 0 {
// 		return val
// 	}
// 	// break string into runes
// 	r := d
// 	left := val[r:]
// 	right := val[:r]
// 	val = left + right
// 	return val
// }

func Clean(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	counter := 0
	for n > 1 {
		// if n == 1 {
		// 	return counter
		// }
		sq := int(math.Sqrt(float64(n)))
		min := sq * sq
		n = n - min
		counter = Clean(n) + 1
		// if  n == 2 {
		// 	return 1
		// }
	}
	return counter
}

func numbersRightRotate(nums []int, k int) []int {
	if k < 0 || len(nums) == 0 {
		return nums
	}
	// fmt.Printf("nums %p array %p len %d cap %d slice %v\n", &nums, &nums[0], len(nums), cap(nums), nums)
	r := len(nums) - k%len(nums)
	left := nums[r:]
	right := nums[:r]
	nums = append(left, right...)
	// fmt.Printf("nums %p array %p len %d cap %d slice %v\n", &nums, &nums[0], len(nums), cap(nums), nums)
	return nums
}
func numbesrLeftRotate(nums []int, k int) []int {
	if k < 0 || len(nums) == 0 {
		return nums
	}
	// fmt.Printf("nums %p array %p len %d cap %d slice %v\n", &nums, &nums[0], len(nums), cap(nums), nums)
	r := len(nums) - k%len(nums)
	left := nums[r-1:]
	right := nums[:r-1]
	nums = append(left, right...)
	// fmt.Printf("nums %p array %p len %d cap %d slice %v\n", &nums, &nums[0], len(nums), cap(nums), nums)
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("nums %p array %p len %d cap %d slice %v\n", &nums, &nums[0], len(nums), cap(nums), nums)

	nums = numbersRightRotate(nums, 3)
	fmt.Printf("right:>nums %p array %p len %d cap %d slice %v\n", &nums, &nums[0], len(nums), cap(nums), nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	nums = numbesrLeftRotate(nums, 3)
	fmt.Printf("left:>nums %p array %p len %d cap %d slice %v\n", &nums, &nums[0], len(nums), cap(nums), nums)

}

/* Input : s = "GeeksforGeeks"
       d = 2
Output : Left Rotation  : "eksforGeeksGe"
        Right Rotation : "ksGeeksforGee"


Input : s = "qwertyu"
       d = 2
Output : Left rotation : "ertyuqw"
        Right rotation : "yuqwert" */
