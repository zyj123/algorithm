package quicksort

func sortInts(nums []int) {
	doSort(nums, 0, len(nums)-1)
}

func doSort(nums []int, low int, high int) {
	if low >= high {
		return
	}
	pivot := nums[high]
	i := low - 1
	for j := low; j <= high; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	doSort(nums, low, i-1)
	doSort(nums, i+1, high)
}
