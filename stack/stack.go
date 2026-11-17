package stack

type Stack struct {
	Items []int
}

func NewStack() *Stack {
	return &Stack{Items: []int{}}
}

func (s *Stack) Push(val int) {
	s.Items = append([]int{val}, s.Items...)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.Items) == 0 {
		return 0, false
	}
	val := s.Items[0]
	s.Items = s.Items[1:]
	return val, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack) IsSorted() bool {
	for i := 0; i < len(s.Items)-1; i++ {
		if s.Items[i] > s.Items[i+1] {
			return false
		}
	}
	return true
}
