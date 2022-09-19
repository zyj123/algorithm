package internal

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	var (
		p1 = m - 1
		p2 = n - 1
	)
	for i := len(nums1) - 1; i >= 0; i-- {
		if p2 < 0 {
			return
		}
		if p1 < 0 {
			copy(nums1[:i+1], nums2[:p2+1])
			return
		}
		if nums1[p1] >= nums2[p2] {
			nums1[i] = nums1[p1]
			p1--
		} else {
			nums1[i] = nums2[p2]
			p2--
		}
	}
}
