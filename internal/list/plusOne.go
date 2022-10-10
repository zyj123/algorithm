package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func plusOne(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	var help func(node *ListNode) int
	help = func(node *ListNode) int {
		if node == nil {
			return 1
		}
		nextAns := help(node.Next)
		val := (node.Val + nextAns) % 10
		ans := (node.Val + nextAns) / 10
		node.Val = val
		return ans
	}
	help(dummyHead)
	if dummyHead.Val == 0 {
		return dummyHead.Next
	}
	return dummyHead
}
