package main

import (
	"container/heap"
	"fmt"
)

// IntHeap is a slice of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push inserts element at it's appropriate position in the heap.
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes minimum element from the heap.
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum
// and removes them in order of priority.
func main() {
	h := &IntHeap{2, 1, 4, 9}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("Minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("Popping: %d\n", heap.Pop(h))
	}
}
