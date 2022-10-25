package main

import "fmt"

type fn func(int) int

func mul(a int) fn {

	return func(b int) int {
		a = a * b
		return a
	}

}

func main() {
	// you can write to stdout for debugging purposes, e.g.
	fmt.Println("This is a debug message")
	a := mul(1)(2)

	fmt.Println(a)
}
