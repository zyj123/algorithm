package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	var (
		xDep int
		yDep int
		xPar *TreeNode
		yPar *TreeNode
	)
	var dfs func(node *TreeNode, depth int, par *TreeNode)
	dfs = func(node *TreeNode, depth int, par *TreeNode) {
		if node == nil {
			return
		}
		if node.Val == x {
			xDep = depth
			xPar = par
		}
		if node.Val == y {
			yDep = depth
			yPar = par
		}
		dfs(node.Left, depth+1, node)
		dfs(node.Right, depth+1, node)
	}
	dfs(root, 0, root)
	if xDep == yDep && xPar != yPar {
		return true
	}
	return false
}
