package binart_search_recursion

import "fmt"

func bs(nums []int, target int) int {
	ans := binarySearch(nums, 0, len(nums)-1, target)
	fmt.Printf("nums: %v target:%d ans:%d", nums, target, ans)
	return ans
}

func binarySearch(nums []int, l, r, target int) int {
	if l > r {
		return -1
	}
	mid := (l + r) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return binarySearch(nums, l, r-1, target)
	}
	if nums[mid] < target {
		return binarySearch(nums, l+1, r, target)
	}
	return -1
}
