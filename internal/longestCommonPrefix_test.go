package internal

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		wantRet string
	}{
		{"t1", args{[]string{"ab", "a"}}, "a"},
		{"t2", args{[]string{"a"}}, "a"},
		{"t3", args{[]string{"flower", "flow", "flight"}}, "fl"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := longestCommonPrefix(tt.args.strs); gotRet != tt.wantRet {
				t.Errorf("longestCommonPrefix() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
