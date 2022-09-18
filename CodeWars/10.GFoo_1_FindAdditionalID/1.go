package main

import (
	"fmt"
)

func Diff(a, b []int) int {
	//b is small list add every item into a map
	m := make(map[int]struct{})
	for _, v := range b {
		//if value already in map skip
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
	}

	//checking bigger list element if they are in map
	for _, v := range a {
		if _, ok := m[v]; !ok {
			return v
		}
	}
	return 0
}

func Find(l1, l2 []int) int {
	if len(l1) > len(l2) {
		return Diff(l1, l2)
	}
	return Diff(l2, l1)
}

func main() {
	a := []int{13, 5, 6, 2, 5}
	b := []int{5, 2, 5, 13}
	r := Find(a, b)
	fmt.Println(r)

	r = Find([]int{14, 27, 1, 4, 2, 50, 3, 1}, []int{2, 4, -4, 3, 1, 1, 14, 27, 50})
	fmt.Println(r)
}
