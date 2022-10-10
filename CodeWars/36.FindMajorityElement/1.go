package main

import "fmt"

func find(arr []int) int {
	m := make(map[int]int)
	mid := len(arr) / 2
	for i := 0; i < len(arr); i++ {
		m[arr[i]] += 1
		// fmt.Printf("k=%v, v=%v \n", i, arr[i])
		if m[arr[i]] > mid {
			return arr[i]
		}
	}
	fmt.Println(m)
	return 0
}

func main() {
	fmt.Println(find([]int{3, 2, 3}))
	fmt.Println(find([]int{2, 2, 1, 1, 1, 2, 2}))
}
