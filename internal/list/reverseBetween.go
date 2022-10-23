package list

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func buildListNode() {
	nums := []int{1, 2, 3, 4, 5}
	pre := &ListNode{}
	cur := pre
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	head := pre.Next

	h := reverseBetween(head, 1, 5)
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var (
		p1 = &ListNode{Next: head}
	)
	for i := 1; i < left; i++ {
		p1 = p1.Next
	}
	var (
		pre     *ListNode
		cur     = p1.Next
		subHead = cur
	)
	for i := 0; i <= (right - left); i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	p1.Next = pre
	subHead.Next = cur
	if left == 1 {
		return p1.Next
	}
	return head
}
