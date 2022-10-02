package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST2(root *TreeNode, low int, high int) int {
	var ret int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if node.Val >= low && node.Val <= high {
			ret += node.Val
		}
		if node.Val > high {
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return ret
}

func rangeSumBST3(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	sum := 0
	sum += rangeSumBST(root.Left, low, high)
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}
	if root.Val > high {
		return sum
	}
	sum += rangeSumBST(root.Right, low, high)
	return sum
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	return root.Val + rangeSumBST3(root.Left, low, high) + rangeSumBST2(root.Right, low, high)
}
