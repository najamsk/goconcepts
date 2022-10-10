package main

import "fmt"

func major(arr []int, k int) bool {
	mid := len(arr) / 2

	i := 0
	for i < len(arr) {
		if arr[i] == k {
			break
		}
		i++
	}
	fmt.Println(i, mid, mid+i)
	if mid+i <= len(arr)-1 && arr[mid+i] == k {
		return true
	}

	return false

}

func main() {

	fmt.Println(major([]int{1, 2, 2}, 2))
	fmt.Println(major([]int{1, 2, 2, 2}, 2))
	fmt.Println(major([]int{1, 1, 2, 2}, 2))

	fmt.Println(major([]int{1, 2, 3, 3, 3, 3, 10}, 3))
	fmt.Println(major([]int{1, 1, 2, 4, 4, 4, 6, 6}, 4))
	fmt.Println(major([]int{1, 1, 1, 2}, 1))
}
