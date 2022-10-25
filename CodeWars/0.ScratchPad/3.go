package main

import "fmt"

func rev(s string) string {
	r := ""
	for _, v := range s {
		r = string(v) + r
	}
	return r
}
func check(s string) bool {
	r := rev(s)
	if r == s {
		return true
	}
	return false

}

func main() {
	fmt.Println(check("najam"))
	fmt.Println(check("aba"))
}
