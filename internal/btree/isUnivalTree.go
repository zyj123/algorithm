package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	var dfs func(node *TreeNode, val int) bool
	dfs = func(node *TreeNode, val int) bool {
		if node == nil {
			return true
		}
		return node.Val == val && dfs(node.Left, val) && dfs(node.Right, val)
	}
	return dfs(root, root.Val)
}
