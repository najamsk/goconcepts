package main

import (
	"fmt"
)

func test(s string) int32 {

	m := make(map[string]int)

	for _, c := range s {
		char := string(c)
		// fmt.Println("c=", char)
		m[char] += 1
	}

	for k, v := range s {
		char := string(v)
		if m[char] == 1 {
			return int32(k + 1)
		}
	}
	return -1

}
func main() {
	fmt.Println(test("hackthegame"))
}
