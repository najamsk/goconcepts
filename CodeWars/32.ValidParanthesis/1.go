package main

import "fmt"

func check(s string) bool {
	// m := make(map[byte]byte, 0, len(s))
	m := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := make([]byte, 0, len(s))

	for k := range s {
		c := s[k]
		//if c is in map add it to stack
		if _, ok := m[c]; ok {
			stack = append(stack, c)
		} else if len(stack) == 0 {
			// starting with closing brackets
			return false
		} else {
			//pop from stack
			last := len(stack) - 1
			e := stack[last]
			stack = stack[:last]
			// poped item value:closing bracket != c:closing then its not valid
			if m[e] != c {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(check("[{()}]"))
}
