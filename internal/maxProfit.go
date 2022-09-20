package internal

func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 0
	var ret int
	for i := 1; i < n; i++ {
		dp[i] = max(0, dp[i-1]+(prices[i]-prices[i-1]))
		ret = max(ret, dp[i])
	}
	return ret
}

//func max(i, j int) int {
//	if i > j {
//		return i
//	}
//	return j
//}
