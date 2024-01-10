package validParentheses

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	//m := map[byte]byte{
	//	')': '(',
	//	']': '[',
	//	'}': '{',
	//}
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == '(' || b == '[' || b == '{' {
			stack.Push(b)
			continue
		}

		if stack.Len() == 0 || !(stack.PopV() == b-1 || stack.PopV() == b-2) {
			return false
		} else {
			stack.Pop()
		}
	}

	return stack.Len() == 0

}

type Stack struct {
	d []byte
	p int
}

func NewStack() *Stack {
	return &Stack{
		d: []byte{},
		p: -1,
	}
}

func (s *Stack) Push(v byte) {
	s.d = append(s.d, v)
	s.p++
}

func (s *Stack) Pop() byte {
	data := s.d[s.p]
	s.d = s.d[0:s.p]
	s.p--
	return data
}

func (s *Stack) PopV() byte {
	return s.d[s.p]
}

func (s *Stack) Len() int {
	return s.p + 1
}

//leetcode submit region end(Prohibit modification and deletion)
