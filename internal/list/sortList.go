package list

import "sort"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	var arr []*ListNode
	for cur := head; cur != nil; cur = cur.Next {
		arr = append(arr, cur)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Val < arr[j].Val
	})
	dummyHead := &ListNode{}
	cur := dummyHead
	for _, node := range arr {
		cur.Next = node
		cur = cur.Next
	}
	return dummyHead.Next
}
