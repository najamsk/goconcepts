package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}
type linklist struct {
	nodes   []node
	current *node
	first   *node
	last    *node
}

func (h *linklist) append(val int) {
	n := node{data: val, next: nil}

	if h.first == nil {
		h.first = &n
		h.current = &n
		h.last = &n
	}

	h.last.next = &n
	h.last = &n
}

func Merge(list1 linklist, list2 linklist) linklist {
	res := linklist{}

	if list1.first == nil {
		return list2
	}

	if list2.first == nil {
		return list1
	}

	for list1.current != nil {
		if list2.current == nil {
			break
		}
		if list1.current.data > list2.current.data {
			res.append(list2.current.data)
			res.append(list1.current.data)
		} else {
			res.append(list1.current.data)
			res.append(list2.current.data)
		}
		//move lists current
		list1.current = list1.current.next
		list2.current = list2.current.next
	}

	if list1.current == nil {
		for list2.current != nil {
			res.append(list2.current.data)
			list2.current = list2.current.next
		}
	}

	if list2.current == nil {
		for list1.current != nil {
			res.append(list1.current.data)
			list1.current = list1.current.next
		}
	}

	return res
}

func main() {
	list := linklist{}
	list.append(2)
	list.append(5)
	list.append(7)
	list.append(8)

	list2 := linklist{}
	list2.append(3)
	list2.append(4)
	list2.append(5)
	list2.append(9)
	list2.append(12)

	list3 := linklist{}
	list4 := linklist{}

	fmt.Println("\n Merge")
	//resetting list currents to first so we can start from begining
	list3.current = list3.first
	list2.current = list2.first
	list4.current = list4.first
	list.current = list.first

	res := Merge(list2, list)
	for res.current != nil {
		fmt.Printf("%#v, ", res.current.data)
		res.current = res.current.next
	}
	fmt.Println("")
	res2 := Merge(list, list2)
	for res2.current != nil {
		fmt.Printf("%#v, ", res2.current.data)
		res2.current = res2.current.next
	}

}
