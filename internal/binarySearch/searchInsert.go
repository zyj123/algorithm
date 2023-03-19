package binarySearch

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for left <= right {
		mid := (left + right) >> 1
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
