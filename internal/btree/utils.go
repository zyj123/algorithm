package btree

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
