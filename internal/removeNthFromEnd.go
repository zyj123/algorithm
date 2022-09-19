package internal

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		preHead = &ListNode{Next: head}
		slow    = preHead
		fast    = preHead
	)
	for ; n >= 0; n-- {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return preHead.Next
}
