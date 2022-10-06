package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func _increasingBST(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) (*TreeNode, *TreeNode)

	dfs = func(node *TreeNode) (*TreeNode, *TreeNode) {
		if node == nil {
			return nil, nil
		}
		var (
			head *TreeNode
			tail *TreeNode
		)
		lHead, lTail := dfs(node.Left)
		rHead, rTail := dfs(node.Right)
		if lHead == nil {
			head = node
		} else {
			head = lHead
			lTail.Right = node
		}
		if rHead == nil {
			tail = node
		} else {
			node.Right = rHead
			tail = rTail
		}
		node.Left = nil
		return head, tail
	}
	head, _ := dfs(root)
	return head
}

func increasingBST(root *TreeNode) *TreeNode {
	var (
		preHead = &TreeNode{}
		curNode = preHead
	)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		curNode.Right = node
		curNode = curNode.Right
		curNode.Left = nil
		inorder(node.Right)
	}
	inorder(root)
	return preHead.Right
}
