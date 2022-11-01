package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	fmt.Println("->h-popping:", h)
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func KlargestElement(nums []int, k int) int {

	h := &IntHeap{}
	for _, v := range nums {
		heap.Push(h, v)
		if len(*h) > k {
			heap.Pop(h)
		}
	}
	return heap.Pop(h).(int)
}

func main() {
	fmt.Println("hi")
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Println("MinHeap before pop:", h)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(h))
	}
	//find k largest in this
	hh := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(KlargestElement(hh, 2))
}
