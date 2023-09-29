package asteroid_collision

type Stack []int

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	var top int
	n := len(*s)
	top, *s = (*s)[n-1], (*s)[:n-1]
	return top
}

func (s *Stack) Top() int {
	n := len(*s)
	return (*s)[n-1]
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func asteroidCollision(asteroids []int) []int {
	s := NewStack()
	for _, asteroid := range asteroids {
		if s.Empty() || !canCollision(s.Top(), asteroid) {
			s.Push(asteroid)
			continue
		}
		for !s.Empty() && canCollision(s.Top(), asteroid) {
			asteroid = collision(s.Pop(), asteroid)
		}
		if asteroid != 0 {
			s.Push(asteroid)
		}
	}
	return *s
}

func canCollision(l, r int) bool {
	return l > 0 && r < 0
}

func collision(a, b int) int {
	if abs(a) == abs(b) {
		return 0
	}
	if abs(a) > abs(b) {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
