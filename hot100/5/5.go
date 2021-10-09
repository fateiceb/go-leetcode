package main

func longestPalindrome(s string) string {
	Lens := len(s)
	if Lens < 2 {
		return s
	}
	//dp法
	//两个维度分别表示s[i:j]中的i,j
	//dp[i][j]表示是否为i到j是否为回文串
	dp := make([][]bool, Lens)
	for i := range dp {
		dp[i] = make([]bool, Lens)
		//长度为一均为
		dp[i][i] = true
	}
	var ans string
	maxLen := 0
	//枚举长度
	for L := 2; L < Lens; L++ {
		//枚举左边界
		for i := 0; i < Lens; i++ {
			j := L + i - 1
			if j >= Lens {
				break
			}
			dp[i][j] = (L == 1 || L == 2 || dp[i+1][j-1]) && s[i] == s[j]
			if dp[i][j] && L > maxLen {
				ans = s[i-1 : j+1]
			}
		}
	}
	return ans

}
