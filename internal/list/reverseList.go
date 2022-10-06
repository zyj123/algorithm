package list

func reverseList(head *ListNode) *ListNode {
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
