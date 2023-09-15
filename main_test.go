package main

import (
	"fmt"
	"testing"
)

func Test_sort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{[]int{-2, 0, -3, -10, 1, 2, 3, 4, 5}}},
		{"t2", args{[]int{2, 1, 3, 100, 2000, 999, 4, 5}}},
		{"t3", args{[]int{5, 4, 3, 2, 1}}},
		{"t4", args{[]int{2, 3, 1}}},
		{"t5", args{[]int{2, 1}}},
		{"t6", args{[]int{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort(tt.args.nums)
			fmt.Printf("after sort: %v\n", tt.args.nums)
		})
	}
}
