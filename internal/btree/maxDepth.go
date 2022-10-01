package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func _maxDepth(root *TreeNode) int {
	var ret int

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			if depth > ret {
				ret = depth
			}
			return
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0)
	return ret
}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := maxDepth(node.Left)
	right := maxDepth(node.Right)
	return max(left, right) + 1
}

func maxDepth2(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(maxDepth(node.Left)+1, maxDepth(node.Right)+1)
}

//func max(i, j int) int {
//	if i > j {
//		return i
//	}
//	return j
//}

func __maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ret int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
		ret++
	}
	return ret
}
