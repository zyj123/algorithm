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
func maxPathSum(root *TreeNode) int {
	ret := math.MinInt
	var pathSum func(node *TreeNode) int
	pathSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lSum := pathSum(node.Left)
		rSum := pathSum(node.Right)
		sum := node.Val
		if lSum > 0 {
			sum += lSum
		}
		if rSum > 0 {
			sum += rSum
		}
		if sum > ret {
			ret = sum
		}
		return max(max(lSum, rSum), 0) + node.Val
	}
	pathSum(root)
	return ret
}
