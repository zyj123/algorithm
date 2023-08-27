package quicksort

import (
	"reflect"
	"testing"
)

func Test_sortInt(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{"t1", []int{3, 1, 2, 4}, []int{1, 2, 3, 4}},
		{"t1", []int{3, 1, 2}, []int{1, 2, 3}},
		{"t1", []int{2, 1}, []int{1, 2}},
		{"t1", []int{1}, []int{1}},
		{"t1", []int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortInts(tt.args)
			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("sortInts() = %v want %v", tt.args, tt.want)
			}
		})
	}
}
