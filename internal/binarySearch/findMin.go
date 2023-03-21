package binarySearch

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[r] > nums[l] {
			return nums[l]
		}
		m := l + (r-l)/2
		if nums[m] < nums[r] {
			r = m
		}
		if nums[m] > nums[r] {
			l = m + 1
		}
	}
	return nums[l]
}
