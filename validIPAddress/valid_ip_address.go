package validIPAddress

import (
	"strings"
)

func validIPAddress(IP string) string {
	if strings.Contains(IP, ".") {
		if validIPv4(IP) {
			return "IPv4"
		}
	}
	if strings.Contains(IP, ":") {
		if validIPv46(IP) {
			return "IPv6"
		}
	}
	return "Neither"
}

func validIPv4(IP string) bool {
	seq := strings.Split(IP, ".")
	if len(seq) != 4 {
		return false
	}
	for _, v := range seq {
		if len(v) == 0 {
			return false
		}
		if len(v) > 1 && v[0] == '0' {
			return false
		}
		i := 0
		for _, b := range v {
			if b < '0' || b > '9' {
				return false
			}
			i = i*10 + int(b-'0')
		}
		if !(i >= 0 && i <= 255) {
			return false
		}
	}
	return true
}

func validIPv46(IP string) bool {
	seq := strings.Split(IP, ":")
	if len(seq) != 8 {
		return false
	}
	for _, v := range seq {
		if !(len(v) >= 1 && len(v) <= 4) {
			return false
		}
		for _, b := range v {
			if !((b >= '0' && b <= '9') || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')) {
				return false
			}
		}
	}
	return true
}
