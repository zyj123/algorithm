package internal

func maxLengthBetweenEqualCharacters(s string) int {
	m := map[byte]int{}
	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		if p, ok := m[s[i]]; !ok {
			m[s[i]] = i
		} else {
			if i-p > r-l {
				l, r = p, i
			}
		}
	}
	ret := (r - l + 1) - 2
	if ret >= 0 {
		return ret
	}
	return -1
}
