package btree

import "testing"

func Test_binaryAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{1, 1}, 3},
		{"t1", args{0, 1}, 1},
		{"t1", args{1, 0}, 2},
		{"t1", args{0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryAdd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("binaryAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
