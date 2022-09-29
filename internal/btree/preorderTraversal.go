package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorderTraversal(root *TreeNode) []int {
	var ret []int
	if root == nil {
		return ret
	}
	s := &stack{}
	s.Push(root)
	for !s.isEmpty() {
		node := s.Pop()
		ret = append(ret, node.Val)
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}
	return ret
}

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

func preorderTraversal2(root *TreeNode) []int {
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
