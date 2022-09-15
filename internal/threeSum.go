package internal

import "sort"

func threeSum(nums []int) [][]int {
	var ret [][]int
	if len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		var (
			l      = i + 1
			r      = len(nums) - 1
			target = -nums[i]
		)
		for l < r {
			if nums[l]+nums[r] == target {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
			if nums[l]+nums[r] < target {
				l++
			}
			if nums[l]+nums[r] > target {
				r--
			}
		}
	}
	return ret
}
