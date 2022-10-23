package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	var (
		slow = head
		fast = head
	)
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var (
		l1 = head
		l2 = slow.Next
	)
	slow.Next = nil
	var (
		pre *ListNode
		cur = l2
	)
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	l2 = pre

	dummyHead := &ListNode{}
	cur = dummyHead
	for l1 != nil || l2 != nil {
		if l1 != nil {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		}
		if l2 != nil {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}
}
