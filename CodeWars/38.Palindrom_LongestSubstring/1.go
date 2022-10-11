package main

import "fmt"

func count(s string) string {
	// res := 0
	rs := ""
	rsLen := 0

	for i := 0; i < len(s); i++ {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > rsLen {
				rs = s[l : r+1]
				rsLen = r - l + 1
			}
			l -= 1
			r += 1
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > rsLen {
				rs = s[l : r+1]
				rsLen = r - l + 1
			}
			l -= 1
			r += 1
		}
	}
	return rs
}

func main() {
	// https://leetcode.com/problems/longest-palindromic-substring/description/
	fmt.Println(count("aaab"))
	// fmt.Println(count("abc"))
	fmt.Println(count("aaa"))
	fmt.Println(count("babad"))
}
