//go中的优先队列
package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *MaxHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}
func (h *MaxHeap) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	maxQ := new(MaxHeap)
	heap.Push(maxQ, 10)
	heap.Push(maxQ, 2)
	heap.Push(maxQ, 22)
	heap.Push(maxQ, 3)
	num := heap.Pop(maxQ).(int)
	fmt.Println(maxQ)
	fmt.Println(num)
	fmt.Println(maxQ)
}
