package gcd

import "testing"

func Test_gcd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{x: 5, y: 3}, 1},
		{"t2", args{x: 4, y: 2}, 2},
		{"t3", args{x: 10, y: 3}, 1},
		{"t4", args{x: 18, y: 2}, 2},
		{"t5", args{x: 15, y: 10}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstDigit(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{x: 10}, 1},
		{"t2", args{x: 1}, 1},
		{"t3", args{x: 123}, 1},
		{"t4", args{x: 255}, 2},
		{"t5", args{x: 33}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstDigit(tt.args.x); got != tt.want {
				t.Errorf("firstDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lastDigit(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{x: 10}, 0},
		{"t2", args{x: 1}, 1},
		{"t3", args{x: 123}, 3},
		{"t4", args{x: 255}, 5},
		{"t5", args{x: 33}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastDigit(tt.args.x); got != tt.want {
				t.Errorf("lastDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
