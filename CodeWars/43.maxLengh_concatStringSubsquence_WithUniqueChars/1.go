package main

import (
	"fmt"
	"math"
)

func overlap(charSet map[string]int, s string) bool {
	prev := map[string]int{}
	r := false
	for _, c := range s {
		_, okChar := charSet[string(c)]
		_, okP := prev[string(c)]
		if okChar || okP {
			return true
		}
		prev[string(c)]++
	}
	return r
}
func back(i int, arr []string, m map[string]int) int {
	res := 0
	if len(arr) == i {
		return len(m)
	}
	if overlap(m, arr[i]) == false {
		for _, c := range arr[i] {
			m[string(c)]++
		}
		res = back(i+1, arr, m)
		for _, c := range arr[i] {
			delete(m, string(c))
		}
	}
	return int(math.Max(float64(res), float64(back(i+1, arr, m))))
}

func maxLength(arr []string) int {
	m := make(map[string]int)

	//back func
	return back(0, arr, m)
}

func main() {
	//https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/
	fmt.Println(maxLength([]string{"un", "iq", "ue"}))
}
