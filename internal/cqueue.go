package internal

type stack struct {
	data []int
}

func (s *stack) Len() int {
	return len(s.data)
}
func (s *stack) Push(value int) {
	s.data = append(s.data, value)
}
func (s *stack) Pop() int {
	old := s.data
	top := old[len(old)-1]
	s.data = old[:len(old)-1]
	return top
}

type CQueue struct {
	in  stack
	out stack
}

func ConstructorQueue() CQueue {
	return CQueue{
		in:  stack{},
		out: stack{},
	}
}

func (c *CQueue) AppendTail(value int) {
	c.in.Push(value)
}

func (c *CQueue) DeleteHead() int {
	if c.out.Len() > 0 {
		return c.out.Pop()
	}
	if c.in.Len() > 0 {
		c.move()
	}
	if c.out.Len() > 0 {
		return c.out.Pop()
	}
	return -1
}

func (c *CQueue) move() {
	for c.in.Len() > 0 {
		c.out.Push(c.in.Pop())
	}
}
