package dp

import "testing"

func Test_numWays(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{0}, 1},
		{"t2", args{1}, 1},
		{"t3", args{2}, 2},
		{"t4", args{3}, 3},
		{"t5", args{4}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.n); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
