package internal

func maxSubArray(nums []int) int {
	var ret int
	preMax := nums[0]
	ret = preMax
	for i := 1; i < len(nums); i++ {
		curMax := max(nums[i], nums[i]+preMax)
		preMax = curMax
		ret = max(ret, curMax)
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
