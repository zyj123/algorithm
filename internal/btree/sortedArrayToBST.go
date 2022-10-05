package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{Val: nums[0]}
	}
	mid := len(nums) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])
	return node
}
