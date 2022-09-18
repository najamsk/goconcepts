package main

import "fmt"

type Node struct {
	key  int
	next *Node
	prev *Node
}
type List struct {
	head   *Node
	first  *Node
	tail   *Node
	mapper map[int]*Node
}

func (l *List) Insert(k int) {
	n := &Node{key: k}
	if l.first == nil {
		l.head = n
		l.first = n
		l.tail = n
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}
	l.mapper[n.key] = n
}

func (l *List) Search(k int) *Node {
	l.head = l.first
	for l.head != nil {
		if l.head.key == k {
			return l.head
		}
		l.head = l.head.next
	}
	return nil
}

func (l *List) Print() {
	l.head = l.first
	for l.head != nil {
		fmt.Println(l.head.key)
		l.head = l.head.next
	}
}

func (l *List) PrintBack() {
	l.head = l.tail
	for l.head != nil {
		fmt.Println(l.head.key)
		l.head = l.head.prev
	}
}

func (l *List) ReadItem(k int) {
	if val, ok := l.mapper[k]; ok {
		// fmt.Println(val.key)
		//detach node and append to last
		if val.next == nil {
			l.Print()
			return
		}
		if val.prev == nil {
			val.next.prev = nil
			l.first = val.next
		} else {
			val.next.prev = val.prev
			val.prev.next = val.next
		}
		val.next = nil
		val.prev = l.tail

		l.tail.next = val
		l.tail = val

		l.Print()
	}
}

func main() {
	fmt.Println("hello")
	m := make(map[int]*Node)
	l := List{mapper: m}
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Print()
	nn := l.Search(2)
	fmt.Printf("next:%d, prev:%d \n", nn.next.key, nn.prev.key)
	fmt.Println("printing back")
	l.PrintBack()
	fmt.Println("mapper")
	val := l.mapper[2]
	fmt.Printf("key:%d,next:%d, prev:%d \n", val.key, val.next.key, val.prev.key)
	fmt.Println("read item")
	l.ReadItem(2)
	fmt.Println("read item")
	l.ReadItem(1)
	fmt.Println("read item")
	l.ReadItem(1)
}
