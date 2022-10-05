package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {
	var ret [][]int
	var dfs func(node *TreeNode, path []int, sum int)
	dfs = func(node *TreeNode, path []int, sum int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		sum += node.Val
		if node.Left == nil && node.Right == nil {
			if sum == target {
				ret = append(ret, path)
			}
		}
		if node.Left != nil {
			lPath := make([]int, len(path))
			copy(lPath, path)
			dfs(node.Left, lPath, sum)
		}
		if node.Right != nil {
			rPath := make([]int, len(path))
			copy(rPath, path)
			dfs(node.Right, rPath, sum)
		}
	}
	dfs(root, []int{}, 0)
	return ret
}
