package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{13}
	res := 0
	for _, v := range nums {
		res ^= v
	}

	fmt.Println(res ^ 13)
}

func vowelStrings(words []string, left int, right int) int {
	mp := map[uint8]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	ans := 0
	for i := left; i <= right; i++ {
		word := words[i]
		if mp[word[0]] && mp[word[len(word)-1]] {
			ans += 1
		}
	}
	return ans
}

func maxScore(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	res := 0
	pre := 0
	for _, v := range nums {
		pre += v
		if pre > 0 {
			res++
		}
	}
	return res
}

func beautifulSubarrays(nums []int) int64 {
	if len(nums) == 1 && nums[0] == 0 {
		return 1
	}
	xors := map[int]int{-1: 0}
	preXOR := 0
	for k, v := range nums {
		preXOR ^= v
		xors[k] = preXOR
	}
	var res int64
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if xors[j]^xors[i-1] == 0 {
				res++
			}
		}
	}
	return res
}
