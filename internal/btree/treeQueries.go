package btree

func treeQueries(root *TreeNode, queries []int) []int {
	heights := make([]int, 0, len(queries))
	for _, q := range queries {
		heights = append(heights, dfs(root, q)-1)
	}
	return heights
}

func dfs(node *TreeNode, query int) int {
	if node == nil {
		return 0
	}
	if node.Val == query {
		return 0
	}
	return 1 + max(dfs(node.Left, query), dfs(node.Right, query))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
