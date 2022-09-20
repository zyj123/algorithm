package internal

func maxProfit(prices []int) int {
	var (
		preMax int
		ret    int
	)
	for i := 1; i < len(prices); i++ {
		preMax = max(0, preMax+(prices[i]-prices[i-1]))
		ret = max(ret, preMax)
	}
	return ret
}

//func max(i, j int) int {
//	if i > j {
//		return i
//	}
//	return j
//}
