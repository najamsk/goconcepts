package main

import "fmt"

func greet(s string, ff func(string) string) string {
	return ff(s)
}

func closuer() func(string) string {
	t := "dr."
	return func(s string) string {
		return t + s
	}
}

func main() {
	dr := closuer()
	vv := greet("najam awan", dr)
	fmt.Println(vv)
}
