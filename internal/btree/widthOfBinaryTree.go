package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ret int
	root.Val = 1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		width := queue[n-1].Val - queue[0].Val + 1
		if width > ret {
			ret = width
		}
		for i := 0; i < n; i++ {
			node := queue[i]
			if node.Left != nil {
				node.Left.Val = 2 * node.Val
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = 2*node.Val + 1
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
	}
	return ret
}
