package btree

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var ret []int
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		ret = append(ret, node.Val)
		for _, c := range node.Children {
			dfs(c)
		}
	}
	dfs(root)
	return ret
}
