package main

import (
	"fmt"
)

func LongestSub(s string) int {
	// fmt.Println(s)
	m := make(map[string]struct{})
	count := 0
	for _, c := range s {
		// fmt.Printf("k=%d, c=%s \n", k, string(c))
		if _, ok := m[string(c)]; ok {
			if len(m) > count {
				count = len(m)
			}
			m = make(map[string]struct{})
			m[string(c)] = struct{}{}
			continue
		}
		m[string(c)] = struct{}{}
	}
	if len(m) > count {
		count = len(m)
	}
	return count
}

func main() {
	fmt.Println("hello")
	fmt.Println("counter = ", LongestSub("abcabcbb"))
}
