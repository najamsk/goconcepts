package main

import "fmt"

type MinHeap struct {
	array []int
}

func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MinHeap) Extract() int {
	if len(h.array) == 0 {
		return -1
	}

	last := len(h.array) - 1
	extracted := h.array[0]
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	h.heapifyDown(0)

	return extracted
}

func (h *MinHeap) heapifyDown(index int) {
	last := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= last {
		//if left is the only child or greather then right it becomes childToCompare
		if l == last {
			childToCompare = l
		} else if h.array[l] < h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		//if index is smaller then child swap
		if h.array[index] > h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			//found correct place
			return
		}
	}

}

func (h *MinHeap) heapifyUp(index int) {
	for h.array[parent(index)] > h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) swap(i1, i2 int) {
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

func main() {
	fmt.Println("hi")
	m := MinHeap{}
	bHeap := []int{80, 10, 90, 30, 5, 70, 50}

	for _, v := range bHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	fmt.Println(m.Extract())
	fmt.Println(m)
}
