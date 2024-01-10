package simplifyPath

import (
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func simplifyPath(path string) string {
	list := strings.Split(path, "/")
	stack := NewStack()

	for _, t := range list {
		if t == "" || t == "." {
			continue
		}

		if strings.Contains(t, ".") {
			switch t {
			case ".":
				continue
			case "..":
				stack.Pop()
			default:
				stack.Push(t)

			}
			continue
		}

		stack.Push(t)
	}

	res := []string{}
	for stack.Len() > 0 {
		res = append(res, stack.Pop())
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	return "/" + strings.Join(res, "/")

}

type Stack struct {
	d []string
	p int
}

func NewStack() *Stack {
	return &Stack{
		d: []string{},
		p: -1,
	}
}

func (s *Stack) Push(v string) {
	s.d = append(s.d, v)
	s.p++
}

func (s *Stack) Pop() string {
	if s.p < 0 {
		return ""
	}
	data := s.d[s.p]
	s.d = s.d[0:s.p]
	s.p--
	return data
}

func (s *Stack) Len() int {
	return s.p + 1
}

//leetcode submit region end(Prohibit modification and deletion)
