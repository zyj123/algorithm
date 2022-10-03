package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root}
	var stop bool
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[i]
			if node == nil {
				stop = true
				continue
			}
			if stop {
				return false
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
		queue = queue[n:]
	}
	return true
}
