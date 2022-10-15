package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func _isPalindrome(head *ListNode) bool {
	var arr []int
	for ; head != nil; head = head.Next {
		arr = append(arr, head.Val)
	}
	n := len(arr)
	for i := 0; i <= n/2; i++ {
		if arr[i] != arr[n-1-i] {
			return false
		}
	}
	return true
}
