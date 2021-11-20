package main

import "sort"

func findLHS(nums []int) (ans int) {
	sort.Ints(nums)
	l := 0
	for r := 0; r < len(nums); r++ {
		for l < r && nums[r]-nums[l] > 1 {
			l++
		}
		if nums[r]-nums[l] == 1 {
			ans = max(ans, r-l+1)
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	arr := []int{1, 1, 1, 1}
	findLHS(arr)
}
