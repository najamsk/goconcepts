package main

import (
	"container/heap"
	"fmt"
)

type HeapInt []int

func (h *HeapInt) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *HeapInt) Push(x any) {
	*h = append(*h, x.(int))
}
func (h HeapInt) Len() int { return len(h) }

func (h HeapInt) Less(i, j int) bool { return h[i] > h[j] }

func (h HeapInt) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func lastStone(a []int) int {
	h := &HeapInt{}
	heap.Init(h)
	for _, v := range a {
		heap.Push(h, v)
	}
	// k := 0
	for h.Len() >= 2 {
		// fmt.Println("heap:", heap.Pop(h))
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)
		diff := a - b
		if diff > 0 {
			heap.Push(h, diff)
		}
	}
	if h.Len() == 0 {
		return 0
	}

	return heap.Pop(h).(int)
}

func main() {
	a := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStone(a))
}
