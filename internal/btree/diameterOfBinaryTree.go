package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	var ret int
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)
		diameter := left + right + 1
		if diameter > ret {
			ret = diameter
		}
		if left >= right {
			return left + 1
		}
		return right + 1
	}
	depth(root)
	return ret
}
