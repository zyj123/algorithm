package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func _preorderTraversal(root *TreeNode) []int {
	var ret []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		ret = append(ret, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ret
}

func preorderTraversal(root *TreeNode) []int {
	var ret []int
	s := &stack{}
	for root != nil || !s.isEmpty() {
		if root != nil {
			ret = append(ret, root.Val)
			s.Push(root)
			root = root.Left
			continue
		}
		node := s.Pop()
		root = node.Right
	}
	return ret
}
