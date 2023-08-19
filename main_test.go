package main

import "testing"

func Test_search(t *testing.T) {

	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"t1", []int{5, 6, 7, 1, 2, 3, 4}, 3, 5},
		{"t2", []int{5, 6, 7, 1, 2, 3, 4}, 1, 3},
		{"t3", []int{5, 6, 7, 1, 2, 3, 4}, 5, 0},
		{"t4", []int{5, 6, 7, 1, 2, 3, 4}, 6, 1},
		{"t5", []int{5, 6, 7, 1, 2, 3, 4}, 4, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
