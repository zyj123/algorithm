package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	root := &TreeNode{}
	var (
		root1Left  *TreeNode
		root1Right *TreeNode
		root2Left  *TreeNode
		root2Right *TreeNode
		root1Val   int
		root2Val   int
	)
	if root1 != nil {
		root1Left, root1Right = root1.Left, root1.Right
		root1Val = root1.Val
	}
	if root2 != nil {
		root2Left, root2Right = root2.Left, root2.Right
		root2Val = root2.Val
	}
	root.Val = root1Val + root2Val
	root.Left = mergeTrees(root1Left, root2Left)
	root.Right = mergeTrees(root1Right, root2Right)
	return root
}
