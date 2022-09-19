package internal

func removeDuplicates(nums []int) int {
	p := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[p] {
			continue
		}
		p++
		if p != i {
			nums[p] = nums[i]
		}
	}
	return p + 1
}
