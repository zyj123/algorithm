package btree

func _postorderTraversal(root *TreeNode) []int {
	var ret []int
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		ret = append(ret, node.Val)
	}
	postorder(root)
	return ret
}

func postorderTraversal(root *TreeNode) []int {
	var ret []int
	s := &stack{}
	var (
		cur  = root
		prev *TreeNode
	)
	for cur != nil || !s.isEmpty() {
		if cur != nil {
			s.Push(cur)
			cur = cur.Left
			continue
		}
		node := s.Pop()
		if node.Right == nil || node.Right == prev {
			ret = append(ret, node.Val)
			prev = node
			cur = nil
			continue
		}
		s.Push(node)
		cur = node.Right
	}
	return ret
}
