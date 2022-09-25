package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	var ret int

	var dfs func(node *TreeNode, sum int) int
	dfs = func(node *TreeNode, sum int) int {
		if node == nil {
			return sum
		}
		if node.Left == nil && node.Right == nil {
			return (sum << 1) | node.Val
		}
		
	}

	return ret
}

func binaryAdd(a, b int) int {
	return (a << 1) | b
}
