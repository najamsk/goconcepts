package main

import "fmt"

func count(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res += 1
			l -= 1
			r += 1
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res += 1
			l -= 1
			r += 1
		}
	}
	return res
}

func main() {
	//https://leetcode.com/problems/palindromic-substrings/
	fmt.Println(count("aaab"))
	fmt.Println(count("abc"))
	fmt.Println(count("aaa"))
}
