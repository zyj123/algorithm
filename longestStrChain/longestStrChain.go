package longestStrChain

import "sort"

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	n := len(words)
	dp := make([]int, n)
	dp[0] = 1
	ans := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if isPredecessor(words[j], words[i]) {
				dp[i] = max(dp[i], dp[j]+1)
				ans = max(ans, dp[i])
			}
		}
	}
	return ans
}

func isPredecessor(word1 string, word2 string) bool {
	if len(word2)-len(word1) != 1 {
		return false
	}

	flag := false
	for p1, p2 := 0, 0; p1 < len(word1); {
		if word1[p1] != word2[p2] {
			if !flag {
				p2++
				flag = true
			} else {
				return false
			}
		} else {
			p1++
			p2++
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
