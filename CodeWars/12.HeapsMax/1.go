package main

import (
	"fmt"
)

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		return -1
	}
	last := len(h.array) - 1
	extracted := h.array[0]
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	last := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= last {
		//if left is the only child or greather then right it becomes childToCompare
		if l == last {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		//if index is smaller then child swap
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			//found correct place
			return
		}
	}

}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return (index * 2) + 1
}

func right(index int) int {
	return (index * 2) + 2
}

// lets implement Max Heaps. Max Heap has max value as root node.
// looks like tree but implemnt as array
// parent Index * 2 + 1 = left child index
// parent Index * 2 + 2 = right child index

func main() {
	fmt.Println("hello boys")
	m := MaxHeap{}
	bHeap := []int{10, 20, 30, 70, 50}

	for _, v := range bHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	m.Extract()
	fmt.Println(m)

}
