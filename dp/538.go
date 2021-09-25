package dp

func minDistance(word1 string, word2 string) (cnt int) {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i, v1 := range word1 {
		for j, v2 := range word2 {
			if v1 == v2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return m + n - 2*dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
