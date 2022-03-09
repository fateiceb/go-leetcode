package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	return quickeSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickeSelect(nums []int, l, r, index int) int {
	q := randomPartition(nums, l, r)
	//主元下标与目标一致，返回第k大数字
	if q == index {
		return nums[q]
	} else if q < index {
		return quickeSelect(nums, q+1, r, index)
	}
	return quickeSelect(nums, l, q-1, index)
}

//随机主元下标，并且更换到nums[r]
func randomPartition(nums []int, l, r int) int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Int()%(r-l+1) + l
	nums[i], nums[r] = nums[r], nums[i]
	return partition(nums, l, r)
}
func partition(nums []int, l, r int) int {
	pivot := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	//更换新的主元
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}
