package heap_sort

import "container/heap"

type HeapInt []int

func (h *HeapInt) Len() int {
	return len(*h)
}

func (h *HeapInt) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *HeapInt) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *HeapInt) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *HeapInt) Pop() (x any) {
	n := len(*h)
	x, *h = (*h)[n-1], (*h)[:n-1]
	return
}

func pickGifts(gifts []int, k int) int64 {
	h := HeapInt(gifts)
	heap.Init(&h)
	for ; k > 0; k-- {
		top := heap.Pop(&h).(int)
		top = sqrt(top)
		heap.Push(&h, top)
	}

	var ans int64
	for i := range h {
		ans += int64(h[i])
	}
	return ans
}

func sqrt(x int) int {
	for i := 1; ; i++ {
		if i*i == x {
			return i
		} else if i*i > x {
			return i - 1
		}
	}
}
