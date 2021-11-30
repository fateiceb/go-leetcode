package main

func partition(s string) (ans [][]string) {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			//s从i+1到j-1的子串为回文串的同时s[i]等于s[j]
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
		}
	}
	path := []string{}
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n {
			ans = append(ans, append([]string{}, path...))
		}
		//枚举子串
		for i := idx; i < n; i++ {
			if dp[idx][i] {
				path = append(path, s[idx:i+1])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}
