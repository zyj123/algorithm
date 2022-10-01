package btree

import (
	"fmt"
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var ret []string
	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node.Left == nil && node.Right == nil {
			ret = append(ret, path+strconv.Itoa(node.Val))
			return
		}
		path += fmt.Sprintf("%d->", node.Val)
		if node.Left != nil {
			dfs(node.Left, path)
		}
		if node.Right != nil {
			dfs(node.Right, path)
		}
	}
	dfs(root, "")
	return ret
}
