package asteroid_collision

import (
	"reflect"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{[]int{5, 10, -5}}, []int{5, 10}},
		{"t2", args{[]int{8, -8}}, []int{}},
		{"t3", args{[]int{10, 2, -5}}, []int{10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroidCollision(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("asteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
