package main

import (
	"sort"
)

func main() {

}

func minFallingPathSum(matrix [][]int) int {
	preMins := matrix[0]
	n := len(matrix)
	for row := 1; row < n; row++ {
		curMins := make([]int, n)
		for col := 0; col < n; col++ {
			curMins[col] = matrix[row][col] + min(preMins, []int{col - 1, col, col + 1})
		}
		preMins = curMins
	}
	sort.Ints(preMins)
	return preMins[0]
}

func min(nums []int, indexs []int) int {
	tmp := make([]int, 0)
	for _, index := range indexs {
		if index >= 0 && index < len(nums) {
			tmp = append(tmp, nums[index])
		}
	}
	sort.Ints(tmp)
	return tmp[0]
}
