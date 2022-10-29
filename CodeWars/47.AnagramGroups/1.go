package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := make(map[string][]string, 0)
	for _, v := range strs {
		ss := sortChars(v)
		m[ss] = append(m[ss], v)
	}
	// fmt.Println(m)
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
func sortChars(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i int, j int) bool { return r[i] < r[j] })
	return string(r)
}
func isPalindrom(s string) bool {
	//clearing string
	s = strings.ReplaceAll(strings.ToLower(s), " ", "")
	off := 0
	for k, v := range s {
		if (v < 97 || v > 122) && (v < 48 || v > 57) {
			k = k - off
			left := s[:k]
			right := s[k+1:]
			s = left + right
			off++
		}
	}
	// fmt.Println(s)
	l := 0
	r := len(s) - 1

	for l <= r {
		// fmt.Printf("l:%v, r:%v \n", l, r)
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true

}
func main() {
	// fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(isPalindrom("0p"))
	// fmt.Println(isPalindrom("A man, a plan, a canal: Panama"))
	// fmt.Println(isPalindrom("amma"))
}
