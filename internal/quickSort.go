package internal

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, i, j int) {
	if j <= i {
		return
	}
	l, r := i, j
	pivot := nums[l]
	for l < r {
		if nums[r] < pivot {
			nums[l] = nums[r]
			l++
			nums[r] = nums[l]
		} else {
			r--
		}
	}
	nums[l] = pivot
	quickSort(nums, i, l-1)
	quickSort(nums, l+1, j)
}
