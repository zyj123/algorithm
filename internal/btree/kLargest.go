package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
	var ret int
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Right)
		k--
		if k == 0 {
			ret = node.Val
			return
		}
		inOrder(node.Left)
	}
	inOrder(root)
	return ret
}
