package main

import (
	"fmt"
)

func max(left, right []int) []int {
	if len(left) > len(right) {
		return left
	}
	return right
}
func ms(left, right []int) []int {
	c := []int{}
	// r := max(left, right)
	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			c = append(c, right[0])
			right = right[1:]
		} else {
			c = append(c, left[0])
			left = left[1:]
		}
	}
	if len(left) > 0 {
		c = append(c, left...)
	}
	if len(right) > 0 {
		c = append(c, right...)
	}

	return c
}

func merge(a []int) []int {
	l := len(a)
	if l == 0 {
		return []int{}
	}
	if l == 1 {
		return a
	}

	//cal mid
	mid := int(l / 2)

	left := a[:mid]
	right := a[mid:]

	left = merge(left)
	right = merge(right)

	return ms(left, right)

}

func main() {
	fmt.Println("hello")

	// a := []int{2, 8, 3, 5}
	a := []int{9, 2, 1, 5, 0}
	fmt.Println(merge(a))
	b := []int{2, 8, 3, 5}
	fmt.Println(merge(b))
}
