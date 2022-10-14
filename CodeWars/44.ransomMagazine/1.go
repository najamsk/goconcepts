package main

import "fmt"

func test(ransomNote, magazine string) bool {
	m := make(map[string]int, 0)
	for _, v := range magazine {
		m[string(v)]++
	}
	for _, v := range ransomNote {
		char := string(v)
		m[char]--
		if m[char] < 0 {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println(test("a", "b"))
	fmt.Println(test("aa", "ab"))
	fmt.Println(test("aa", "aab"))
}
