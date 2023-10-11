package circular_queue

type MyCircularQueue struct {
	queue []int
	head  int
	l     int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k),
		head:  0,
		l:     0,
	}
}

func (cq *MyCircularQueue) EnQueue(value int) bool {
	if cq.IsFull() {
		return false
	}
	n, head, l := len(cq.queue), cq.head, cq.l
	pos := head + l
	if pos >= n {
		pos -= n
	}
	cq.queue[pos] = value
	l += 1
	cq.l = l
	return true
}

func (cq *MyCircularQueue) DeQueue() bool {
	if cq.IsEmpty() {
		return false
	}
	n, head, l := len(cq.queue), cq.head, cq.l
	head += 1
	if head >= n {
		head = 0
	}
	l -= 1

	cq.head, cq.l = head, l
	return true
}

func (cq *MyCircularQueue) Front() int {
	if cq.IsEmpty() {
		return -1
	}
	return cq.queue[cq.head]
}

func (cq *MyCircularQueue) Rear() int {
	if cq.IsEmpty() {
		return -1
	}
	n, head, l := len(cq.queue), cq.head, cq.l
	pos := head + l - 1
	if pos >= n {
		pos -= n
	}
	return cq.queue[pos]
}

func (cq *MyCircularQueue) IsEmpty() bool {
	return cq.l == 0
}

func (cq *MyCircularQueue) IsFull() bool {
	return len(cq.queue) == cq.l
}
