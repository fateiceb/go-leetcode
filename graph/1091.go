package main

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	//左上角或右下角不可访问
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	//dfs爆破
	dy := []int{-1, 1, 0, 0, 1, -1, -1, 1}
	dx := []int{0, 0, 1, -1, 1, -1, 1, -1}
	ans := 0
	var dfs func(r, c, res int)
	dfs = func(r, c, res int) {
		if r < 0 || c < 0 || r >= n || c >= n || grid[r][c] == 1 {
			return
		}
		if r == n-1 && c == n-1 {
			ans = min(ans, res)
		}
		if grid[r][c] == 0 {
			grid[r][c] = 1
			res++
		}
		for i := 0; i < len(dy); i++ {
			dfs(r+dx[i], c+dy[i], res)
		}

	}
	dfs(0, 0, 0)
	return ans
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	testcase := [][]int{{0, 1}, {1, 0}}
	shortestPathBinaryMatrix(testcase)
}
