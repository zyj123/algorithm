package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		if !isSameTree(p.Left, q.Left) {
			return false
		}
		return isSameTree(p.Right, q.Right)
	}
	return false
}
