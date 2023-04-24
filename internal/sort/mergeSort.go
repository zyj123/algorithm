package sort

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	m := len(nums) / 2
	left := mergeSort(nums[:m])
	right := mergeSort(nums[m:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	ret := make([]int, 0, len(left)+len(right))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			ret = append(ret, left[l])
			l++
		} else {
			ret = append(ret, right[r])
			r++
		}
	}
	if l < len(left) {
		ret = append(ret, left[l:]...)
	}
	if r < len(right) {
		ret = append(ret, right[r:]...)
	}

	return ret
}
