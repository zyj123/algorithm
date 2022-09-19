package internal

import (
	"reflect"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want [][]int
	}{
		{args{nums: []int{1, 2, 3}}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := permuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
