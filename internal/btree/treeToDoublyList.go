package btree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var (
		head = &TreeNode{}
		cur  = head
	)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		{
			cur.Right = node
			node.Left = cur
			cur = cur.Right
		}
		inorder(node.Right)
	}
	inorder(root)
	cur.Right = head.Right
	cur.Right.Left = cur
	return head.Right
}
