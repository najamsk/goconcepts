package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

func Invert(root *node) *node {
	if root == nil {
		return root
	}
	// fmt.Println(root.data)
	left := Invert(root.left)
	right := Invert(root.right)
	root.left = right
	root.right = left
	return root
}

func PrintTree(root *node) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	PrintTree(root.left)
	PrintTree(root.right)
}

func main() {
	fmt.Println("hi")
	n1 := &node{data: 1}
	n3 := &node{data: 3}
	n6 := &node{data: 6}
	n9 := &node{data: 9}

	n2 := &node{data: 2, left: n1, right: n3}
	n7 := &node{data: 7, left: n6, right: n9}
	n4 := &node{data: 4, left: n2, right: n7}

	PrintTree(n4)
	fmt.Println("==========")
	Invert(n4)
	PrintTree(n4)
}
