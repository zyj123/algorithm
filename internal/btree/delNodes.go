package btree

var forests []*TreeNode
var toDelete map[int]struct{}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	forests = []*TreeNode{}
	toDelete = make(map[int]struct{})
	for _, v := range to_delete {
		toDelete[v] = struct{}{}
	}
	dfs(root, true)
	return forests
}

func dfs(node *TreeNode, isRoot bool) *TreeNode {
	if node == nil {
		return nil
	}

	if _, ok := toDelete[node.Val]; ok {
		dfs(node.Left, true)
		dfs(node.Right, true)
		return nil
	}
	node.Left = dfs(node.Left, false)
	node.Right = dfs(node.Right, false)
	if isRoot {
		forests = append(forests, node)
	}
	return node
}
