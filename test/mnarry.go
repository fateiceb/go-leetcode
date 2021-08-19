package main

import "fmt"

func main() {
	fmt.Println(test([]int{2, 2}, 2, 50)) //expect 4
}
func test(arr []int, m, k int) int {
	mod := int(1e9 + 7)
	n := len(arr)
	ans := 0
	var backtrack func(idx int)
	path := make([]int, 0)
	backtrack = func(idx int) {
		if len(path) == n && dis(path, arr) <= k {
			ans = (ans + 1) % mod
			return
		}
		if idx > n {
			return
		}
		for i := 1; i <= m; i++ {
			path = append(path, i)
			backtrack(idx + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}

func dis(a, b []int) (result int) {

	for i := range b {
		result += abs(a[i] - b[i])
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
