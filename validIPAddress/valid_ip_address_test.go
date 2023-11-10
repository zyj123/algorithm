package validIPAddress

import "testing"

func Test_validIPAddress(t *testing.T) {
	type args struct {
		IP string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"1.0.1."}, "Neither"},
		{"t1", args{"20EE:FGb8:85a3:0:0:8A2E:0370:7334"}, "Neither"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validIPAddress(tt.args.IP); got != tt.want {
				t.Errorf("validIPAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
