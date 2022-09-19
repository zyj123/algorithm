package internal

import (
	"reflect"
	"testing"
)

func exchange(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l]%2 == 0 {
			l++
		}
		for l < r && nums[r]%2 == 1 {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}

func Test_exchange(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exchange(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exchange() = %v, want %v", got, tt.want)
			}
		})
	}
}
