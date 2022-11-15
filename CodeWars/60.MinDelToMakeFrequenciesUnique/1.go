package main

import (
	"fmt"
	"math"
)

func test(s string) int {
	m := make(map[string]int, len(s))
	l := 0

	for _, v := range s {
		f := m[string(v)] + 1
		l = int(math.Max(float64(f), float64(l)))
		m[string(v)] = f
	}

	// fmt.Println(m)
	// fmt.Println(l)
	dels := 0
	for _, v := range m {
		if v > l {
			dels++
		}
		l--
	}

	return dels
}

func main() {
	fmt.Println(test("aabbbcca"))
	fmt.Println(test("aab"))
	fmt.Println(test("ceabaacb"))

}
