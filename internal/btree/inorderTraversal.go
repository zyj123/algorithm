package btree

func _inorderTraversal(root *TreeNode) []int {
	var ret []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ret = append(ret, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return ret
}

type stack []*TreeNode

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Push(node *TreeNode) {
	*s = append(*s, node)
}
func (s *stack) Pop() *TreeNode {
	if len(*s) == 0 {
		return nil
	}
	n := len(*s)
	ret := (*s)[n-1]
	*s = (*s)[:n-1]
	return ret
}

func inorderTraversal2(root *TreeNode) []int {
	var ret []int
	s := stack{}
	for root != nil || !s.isEmpty() {
		for root != nil {
			s.Push(root)
			root = root.Left
		}
		node := s.Pop()
		ret = append(ret, node.Val)
		root = node.Right
	}
	return ret
}
