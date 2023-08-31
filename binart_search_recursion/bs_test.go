package binart_search_recursion

import "testing"

func Test_bs(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1, 2, 3, 4, 5}, 2}, 1},
		{"t1", args{[]int{1, 2, 3, 4, 5}, 1}, 0},
		{"t1", args{[]int{1, 2, 3, 4, 5}, 5}, 4},
		{"t1", args{[]int{1, 2, 3, 4, 5}, 6}, -1},
		{"t1", args{[]int{1, 2}, 1}, 0},
		{"t1", args{[]int{1, 2}, 2}, 1},
		{"t1", args{[]int{1}, 1}, 0},
		{"t1", args{[]int{}, 1}, -1},
		{"t1", args{[]int{1, 2, 4, 5}, 3}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bs(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("bs() = %v, want %v", got, tt.want)
			}
		})
	}
}
