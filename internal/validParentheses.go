package internal

type Stack struct {
	data []byte
}

func (s *Stack) Len() int {
	return len(s.data)
}
func (s *Stack) Push(val byte) {
	s.data = append(s.data, val)
}
func (s *Stack) Pop() byte {
	ret := s.Get()
	s.data = s.data[:len(s.data)-1]
	return ret
}
func (s *Stack) Get() byte {
	return s.data[len(s.data)-1]
}

func isValid(s string) bool {
	stack := &Stack{}
	for _, v := range s {
		if stack.Len() == 0 || !isPair(byte(v), stack.Get()) {
			stack.Push(byte(v))
		} else {
			stack.Pop()
		}
	}
	return stack.Len() == 0
}

func isPair(r, l byte) bool {
	var pair = map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	return pair[l] == r
}

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//每个右括号都有一个对应的相同类型的左括号。
