package main

import "fmt"

type Treenode struct {
	Val   int
	Left  *Treenode
	Right *Treenode
}

func df(r *Treenode, l int, m *int) {
	if r == nil {
		return
	}
	if *m < l {
		*m = l
	}
	df(r.Left, l+1, m)
	df(r.Right, l+1, m)
}

func maxDepth(r *Treenode) int {
	max := 0
	level := 1
	df(r, level, &max)
	return max
}

func main() {
	n15 := &Treenode{Val: 15}
	n7 := &Treenode{Val: 7}
	n9 := &Treenode{Val: 9}

	n20 := &Treenode{Val: 20, Left: n15, Right: n7}
	// n20 := &Treenode{data: 20}
	n3 := &Treenode{Val: 3, Left: n9, Right: n20}

	fmt.Println(maxDepth(n3))

}
