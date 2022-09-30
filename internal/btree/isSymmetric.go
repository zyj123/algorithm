package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isNodeEqual(root.Left, root.Right)
}

func isNodeEqual(nodeA, nodeB *TreeNode) bool {
	if nodeA == nil && nodeB == nil {
		return true
	}
	if nodeA != nil && nodeB != nil && nodeA.Val == nodeB.Val {
		if !isNodeEqual(nodeA.Left, nodeB.Right) {
			return false
		}
		return isNodeEqual(nodeA.Right, nodeB.Left)
	}
	return false
}
