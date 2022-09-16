package internal

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{"()"}, true},
		{args{"{[()]}"}, true},
		{args{"()[]{}"}, true},
		{args{"(]"}, false},
		{args{"}}}{{{"}, false},
	}
	for _, tt := range tests {
		t.Run("t", func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
