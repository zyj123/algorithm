package internal

import "testing"

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{grid: [][]byte{{'1'}}},
			1,
		},
		{
			args{grid: [][]byte{{'0'}}},
			0,
		},
		{
			args{grid: [][]byte{{'1', '0'}, {'0', '1'}}},
			2,
		},
		{
			args{grid: [][]byte{{'1', '1'}, {'0', '1'}}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
