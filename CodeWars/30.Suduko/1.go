package main

import "fmt"

func checkOpts(a []int) []int {
	m := make(map[int]bool, 9)
	for i := 1; i <= 9; i++ {
		m[i] = true
	}
	for _, v := range a {
		if m[v] == true {
			m[v] = false
		}
	}
	fmt.Println(m)
	return a
}

func main() {

	fmt.Println(byte(1))

	fmt.Println(checkOpts([]int{0, 1, 0, 3, 5, 0, 0, 0, 0}))
}
