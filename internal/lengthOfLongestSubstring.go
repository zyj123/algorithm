package internal

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	l := len(s)
	ret := 1
	for i := 0; i < l-1; i++ {
		bucket := map[byte]bool{}
		j := i
		for ; j < l; j++ {
			if bucket[s[j]] {
				break
			}
			bucket[s[j]] = true
		}
		r := j - i
		if r > ret {
			ret = r
		}
	}
	return ret
}
