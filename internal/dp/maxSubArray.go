package dp

func _maxSubArray(nums []int) int {
	var (
		pre = -10 ^ 4
		res = pre
	)
	for _, v := range nums {
		pre = max(v, pre+v)
		res = max(res, pre)
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	var (
		pre = -10 ^ 4
		res = pre
	)
	for _, v := range nums {
		pre = max(v, pre+v)
		res = max(res, pre)
	}
	return res
}
