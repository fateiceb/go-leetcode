package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	expect := 5
	real := findKthLargest(nums, k)
	if real == expect {
		fmt.Println("it's ok")
	}
}
func findKthLargest(nums []int, k int) int {
	heapsize := len(nums)
	buildMaxHeap(nums, heapsize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapsize--
		maxheapify(nums, 0, heapsize)
	}
	return nums[0]
}
func buildMaxHeap(nums []int, heapsize int) {
	for i := heapsize / 2; i >= 0; i-- {
		maxheapify(nums, i, heapsize)
	}
}
func maxheapify(nums []int, i int, heapsize int) {
	l, r := i*2+1, i*2+2
	largest := i
	if l < heapsize && nums[l] > nums[largest] {
		largest = l
	}
	if r < heapsize && nums[r] > nums[largest] {
		largest = r
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		maxheapify(nums, largest, heapsize)
	}
}
