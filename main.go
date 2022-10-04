package main

import "fmt"

func main() {
	ret := [][]int{}
	nums := make([]int, 0, 4)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	ret = append(ret, nums)
	fmt.Println(ret)
	nums = append(nums, 4)
	fmt.Println(ret)
}
