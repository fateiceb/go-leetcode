package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type aheap struct {
	sort.IntSlice
}

func (h *aheap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}
func (h *aheap) Pop() interface{} {
	a := h.IntSlice
	x := a[:len(a)-1]
	h.IntSlice = a
	return x
}
func findKthLargest(nums []int, k int) int {
	h := &aheap{}
	heap.Init(h)
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}
	for i := k; i < len(nums); i++ {
		if nums[i] > h.IntSlice[0] {
			h.IntSlice[0] = nums[i]
			heap.Fix(h, 0)
		}
	}
	return h.IntSlice[0]
}
func main() {
	input1 := []int{3, 2, 1, 5, 6, 4}
	input2 := 2
	ans := findKthLargest(input1, input2)
	fmt.Println(ans)
}
