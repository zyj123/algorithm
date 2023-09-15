package min_number_in_rotate_array

import "testing"

func Test_minArray(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{3, 4, 5, 1, 2}}, 1},
		{"t2", args{[]int{5, 1, 2, 3, 4}}, 1},
		{"t3", args{[]int{1, 2, 3, 4, 5}}, 1},
		{"t4", args{[]int{1, 1, 1, 0, 1, 1}}, 0},
		{"t5", args{[]int{1, 2}}, 1},
		{"t6", args{[]int{2, 1}}, 1},
		{"t7", args{[]int{1}}, 1},
		{"t8", args{[]int{3, 1, 3}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minArray(tt.args.numbers); got != tt.want {
				t.Errorf("minArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
