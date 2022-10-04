package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	ret := true
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if !ret {
			return 0
		}
		if node == nil {
			return 0
		}
		lDep := depth(node.Left)
		rDep := depth(node.Right)
		if abs(lDep-rDep) > 1 {
			ret = false
		}
		return 1 + max(lDep, rDep)
	}
	depth(root)
	return ret
}
