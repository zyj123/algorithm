package btree

func isBalanced(root *TreeNode) bool {
	ret := true
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := 1 + dfs(node.Left)
		right := 1 + dfs(node.Right)
		ret = ret && (abs(left-right) <= 1)
		return max(left, right)
	}
	dfs(root)
	return ret
}
