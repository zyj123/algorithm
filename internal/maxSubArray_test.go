package internal

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{nums: []int{1, 2, 3}}, 6},
		{args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 6},
		{args{nums: []int{1}}, 1},
		{args{nums: []int{5, 4, -1, 7, 8}}, 23},
		{args{nums: []int{0}}, 0},
		{args{nums: []int{-1}}, -1},
	}
	for _, tt := range tests {
		t.Run("t", func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
