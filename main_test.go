package main

import (
	"testing"
)

func Test_sqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{x: 1}, 1},
		{"t1", args{x: 2}, 1},
		{"t1", args{x: 3}, 1},
		{"t1", args{x: 4}, 2},
		{"t1", args{x: 6}, 2},
		{"t1", args{x: 9}, 3},
		{"t1", args{x: 15}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sqrt(tt.args.x); got != tt.want {
				t.Errorf("sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
