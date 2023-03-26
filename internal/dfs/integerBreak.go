package dfs

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i-2; j++ {
			dp[i] = max(dp[i], j*max(dp[i-j], i-j))
		}
	}

	return dp[n]
}
