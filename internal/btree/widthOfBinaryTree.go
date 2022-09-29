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
	var ret int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[i]
			if node == nil {
				continue
			}
			queue = append(queue, []*TreeNode{node.Left, node.Right}...)
		}
		l := n
		if queue[0] == nil {
			l--
		}
		if n != 1 && queue[n-1] == nil {
			l--
		}
		queue = queue[n:]
		if l > ret {
			ret = l
		}
	}
	return ret
}
