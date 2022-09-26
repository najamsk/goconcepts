package main

import "fmt"

func LinearSearch(items []int, target int) int {
	left := 0
	right := len(items) - 1

	for left <= right {
		mid := int((left + right) / 2)
		if items[mid] == target {
			return mid
		}
		if items[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	if left <= right && items[left] == target {
		return left
	}
	return -1
}

func bRecurring(items []int, left, right, target int) int {
	if left > right {
		return -1
	}
	mid := int((left + right) / 2)
	if target == items[mid] {
		return mid
	}
	if items[mid] < target {
		left = mid + 1
	} else {
		right = mid - 1
	}
	return bRecurring(items, left, right, target)
}

func searchRecuring(items []int, target int) int {
	left := 0
	right := len(items) - 1
	return bRecurring(items, left, right, target)
}

func PeekFinder(items []int) int {
	left := 0
	right := len(items) - 1

	for left < right {
		mid := int((left + right) / 2)
		if items[mid] < items[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	fmt.Println(LinearSearch([]int{0, 3, 5, 8, 12, 20}, 8))
	fmt.Println(searchRecuring([]int{0, 3, 5, 8, 12, 20}, 20))
	fmt.Println(searchRecuring([]int{0, 3, 5, 8, 12, 20}, 5))
	fmt.Println(searchRecuring([]int{0, 3, 5, 8, 12, 20}, 0))
	fmt.Println(searchRecuring([]int{0, 3, 5, 8, 12, 20}, 90))
	// fmt.Println(PeekFinder([]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}))
	fmt.Println(PeekFinder([]int{1, 2, 1}))
}
