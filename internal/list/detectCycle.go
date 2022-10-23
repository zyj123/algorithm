package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var (
		slow  = head
		fast  = head
		slow2 = head
	)
	isMove := false
	for slow != fast || !isMove {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		isMove = true
	}
	for slow != slow2 {
		slow = slow.Next
		slow2 = slow2.Next
	}
	return slow
}
