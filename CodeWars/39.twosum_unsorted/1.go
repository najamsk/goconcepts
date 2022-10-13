package main

import "fmt"

func twosum(a []int, t int) []int {
	res := []int{}
	m := make(map[int]int, 0)
	for k, v := range a {
		diff := t - v
		if val, ok := m[diff]; ok {
			res = append(res, k, val)
			return res
		}
		m[v] = k
	}
	return res
}

func main() {
	ar := []int{2, 1, 3, 5}
	fmt.Println(twosum(ar, 4))
}
