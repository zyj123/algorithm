package kth_largest_element

import "container/heap"

type heapInt []int

func (h *heapInt) Len() int {
	return len(*h)
}
func (h *heapInt) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *heapInt) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *heapInt) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}
func (h *heapInt) Push(v any) {
	*h = append(*h, v.(int))
}

func findKthLargest(nums []int, k int) int {
	h := &heapInt{}
	for _, num := range nums {
		*h = append(*h, num)
	}
	heap.Init(h)
	for h.Len() > k {
		heap.Pop(h)
	}
	return heap.Pop(h).(int)
}
