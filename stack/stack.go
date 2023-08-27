package stack

type Stack []int

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}
func (s *Stack) Peek() int {
	n := len(*s)
	return (*s)[n-1]
}
func (s *Stack) Pop() int {
	n := len(*s)
	val := (*s)[n-1]
	*s = (*s)[:n-1]
	return val
}
func (s *Stack) Size() int {
	return len(*s)
}
func (s *Stack) Empty() bool {
	return len(*s) == 0
}
