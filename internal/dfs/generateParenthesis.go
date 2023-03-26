package dfs

func generateParenthesis(n int) []string {
	var pairs []string
	var dfs func(string, int, int)
	dfs = func(pair string, l, r int) {
		if l < 0 || l > r {
			return
		}
		if l == 0 && r == 0 {
			pairs = append(pairs, pair)
			return
		}

		dfs(pair+"(", l-1, r)
		dfs(pair+")", l, r-1)
	}

	dfs("", n, n)
	return pairs
}
