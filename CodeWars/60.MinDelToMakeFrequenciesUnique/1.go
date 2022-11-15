package main

import (
	"fmt"
	"sort"
)

func test(s string) int {

	frequency, res := make([]int, 26), 0
	for i := 0; i < len(s); i++ {
		frequency[s[i]-'a']++
	}
	is := sort.IntSlice(frequency)
	sr := sort.Reverse(is)
	sort.Sort(sr)
	// sort.Sort(sort.Reverse(sort.IntSlice(frequency)))
	for i := 1; i <= 25; i++ {
		if frequency[i] == frequency[i-1] && frequency[i] != 0 {
			res++
			frequency[i]--
			i--
			sort.Sort(sort.Reverse(sort.IntSlice(frequency)))
		}
	}
	return res

}

func main() {
	// fmt.Println(test("aabbbcca"))
	// fmt.Println(test("aab"))

	fmt.Println(test("ceabaacb"))
	// fmt.Println(test("abcabc"))

}
