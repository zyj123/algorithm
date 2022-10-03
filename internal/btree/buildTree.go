package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	node := &TreeNode{}
	node.Val = preorder[0]
	i := findIndex(inorder, node.Val)
	preorder = preorder[1:]
	node.Left = buildTree(preorder[:i], inorder[:i])
	node.Right = buildTree(preorder[i:], inorder[i+1:])
	return node
}

func findIndex(nums []int, val int) int {
	for k, v := range nums {
		if val == v {
			return k
		}
	}
	return -1
}
