package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func _reverseList(head *ListNode) *ListNode {
	var (
		pre  *ListNode = nil
		curr           = head
	)
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	reverse := reverseList(head)
	head.Next.Next = head
	head.Next = nil
	return reverse
}
