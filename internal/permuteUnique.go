package internal

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	visited := make([]bool, n)
	ret := make([][]int, 0)
	item := make([]int, 0, n)
	var bt func()
	bt = func() {
		if len(item) == n {
			cp := make([]int, 0)
			copy(cp, item)
			ret = append(ret, cp)
		}
		for k, num := range nums {
			if visited[k] {
				continue
			}
			item = append(item, num)
			visited[k] = true
			bt()
			item = item[:len(item)-1]
			visited[k] = false
		}
	}
	bt()
	return ret
}
