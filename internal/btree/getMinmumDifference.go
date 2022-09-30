package btree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	var list []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		list = append(list, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	ret := math.MaxInt
	for i := 1; i < len(list); i++ {
		if abs(list[i]-list[i-1]) < ret {
			ret = abs(list[i] - list[i-1])
		}
	}
	return ret
}
