package dp

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	var p1, p2 int
	for ; p2 < len(t); p2++ {
		if t[p2] == s[p1] {
			p1++
		}
		if p1 > len(s)-1 {
			return true
		}
	}
	return false
}
