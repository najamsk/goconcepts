package main

import (
	"fmt"
	"sort"
)

func assign(arr []int) []int {
	l := len(arr)
	m := make(map[int]int, 0)
	//if empty or one element
	if l <= 1 {
		return arr
	}

	fmt.Println("input arr:", arr)
	// t := make([]int, l)
	//copy values to t
	t := append([]int(nil), arr...)
	sort.Ints(t)
	// sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	// fmt.Println("input arr:", arr)
	fmt.Println("sorted t:", t)
	pp := 1
	for k, v := range t {
		if k > 0 {
			if t[k-1] != t[k] {
				pp++
			}
		}
		m[v] = pp
	}
	fmt.Println("m:", m)
	for k, v := range arr {
		fmt.Printf("v:%v, m[v]:%v \n", v, m[v])
		arr[k] = m[v]
	}

	return arr
}

func main() {
	// https://leetcode.com/discuss/interview-question/1362915/amazon-hackerrank-question-priority-assignment
	fmt.Println(assign([]int{1, 4, 8, 4}))
}
