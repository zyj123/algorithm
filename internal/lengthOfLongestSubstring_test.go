package internal

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want int
	}{
		{args{"aa"}, 1},
		{args{"au"}, 2},
		{args{"abca"}, 3},
		{args{"a"}, 1},
		{args{""}, 0},
		{args{"abcabcbb"}, 3},
		{args{"pwwkew"}, 3},
	}
	for _, tt := range tests {
		t.Run("t", func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
