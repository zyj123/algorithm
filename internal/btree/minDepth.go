package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[i]
			if node.Left == nil && node.Right == nil {
				return ret
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
		ret++
	}
	return ret
}
