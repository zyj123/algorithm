package internal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		n := len(q)
		level := make([]int, 0, n)
		for i := 0; i < n; i++ {
			node := q[i]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[n:]
		ret = append(ret, level)
	}
	return ret
}
