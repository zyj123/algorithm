package binarySearch

func searchRange(nums []int, target int) []int {
	var (
		s = 0
		e = 0
	)
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r + 1) / 2
		if nums[m] >= target {
			r = m - 1
		} else {
			l = m
		}
	}
	s = l + 1
	l, r = 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}
	e = l - 1
	if e < s {
		return []int{-1, -1}
	}
	return []int{s, e}
}
