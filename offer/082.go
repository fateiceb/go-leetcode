package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) (ans [][]int) {
	path := make([]int, 0, 16)
	sort.Ints(candidates)
	visit := make([]bool, len(candidates))
	var backtrack func(idx int, remain int)
	backtrack = func(idx int, remain int) {
		if remain == 0 {
			ans = append(ans, append([]int{}, path...))
		}
		if idx > len(candidates) || remain < 0 {
			return
		}
		for i := idx; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}

			if i > 0 && !visit[i-1] && candidates[i] == candidates[i-1] {
				continue
			}
			visit[i] = true
			path = append(path, candidates[i])
			backtrack(i+1, remain-candidates[i])
			path = path[:len(path)-1]
			visit[i] = false
		}

	}
	backtrack(0, target)
	return
}
func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
