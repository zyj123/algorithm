package internal

func moveZeroes(nums []int) {
	var l int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[l] = nums[i]
			l++
		}
	}
	for i := l; i < len(nums); i++ {
		nums[i] = 0
	}
}
