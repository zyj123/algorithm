package internal

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want [][]int
	}{
		{args{[]int{0}}, nil},
		{args{[]int{0, 0, 0}}, [][]int{{0, 0, 0}}},
		{args{[]int{1, 2, 3}}, nil},
		{args{[]int{0, 1, 2}}, nil},
		{args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run("t", func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
