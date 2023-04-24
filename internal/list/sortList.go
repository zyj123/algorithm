package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func mergeSort(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	k1, k2 := head, head.Next
//	for ; k2 != nil && k2.Next != nil; k2 = k2.Next.Next {
//		k1 = k1.Next
//	}
//	right := mergeSort(k1.Next)
//	k1.Next = nil
//	left := mergeSort(head)
//	return merge(left, right)
//}
//
//func merge(left, right *ListNode) *ListNode {
//	var (
//		head = &ListNode{}
//		cur  = head
//	)
//	var (
//		h1 = left
//		h2 = right
//	)
//
//	for h1 != nil && h2 != nil {
//		if h1.Val <= h2.Val {
//			cur.Next = h1
//			h1 = h1.Next
//		} else {
//			cur.Next = h2
//			h2 = h2.Next
//		}
//		cur = cur.Next
//	}
//	if h1 != nil {
//		cur.Next = h1
//	}
//	if h2 != nil {
//		cur.Next = h2
//	}
//	return head.Next
//}
