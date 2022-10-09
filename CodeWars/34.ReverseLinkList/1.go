package main

import "fmt"

type Node struct {
	data int
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func (l *List) Insert(k int) {
	n := &Node{data: k}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}

func (l *List) Reverse() {
	var next *Node
	var prev *Node
	cur := l.Head

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	l.Head = prev
}

func (l *List) Print() {
	c := l.Head
	for c != nil {
		fmt.Println(c.data)
		c = c.Next
	}
}

func main() {
	fmt.Println("hello")
	l := &List{}
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Print()
	fmt.Println("===============")
	l.Reverse()
	l.Print()
}
