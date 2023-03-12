package dp

import "math"

func maxProfit(prices []int) int {
	preMin := math.MaxInt
	res := 0
	for _, price := range prices {
		if price < preMin {
			preMin = price
		}
		if price-preMin > res {
			res = price - preMin
		}
	}
	return res
}
