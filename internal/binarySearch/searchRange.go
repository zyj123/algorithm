package binarySearch

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if target < nums[m] {
			l = m + 1
		} else if target == nums[m] {
			r = m
		} else if target > nums[m] {
			r = m - 1
		}
	}
	if l < len(nums) || nums[l] != target {
		return []int{-1, -1}
	}
	s := l
	l, r = 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if target < nums[m] {
			r = m - 1
		} else if target == nums[m] {
			l = m
		} else if target > nums[m] {
			l = m + 1
		}
	}
	e := l
	return []int{s, e}
}
