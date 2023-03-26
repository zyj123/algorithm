package dfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := -1001

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val
		left := dfs(node.Left)
		right := dfs(node.Right)
		sum += max(left, 0)
		sum += max(right, 0)
		res = max(res, sum)
		return node.Val + max(0, max(left, right))
	}
	dfs(root)
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
