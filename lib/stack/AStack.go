package stack

type AStack struct {
	size      int
	listArray []int
}

func (s *AStack) Clear() {
	s.size = 0
}

func (s *AStack) Push(val int) {
	s.listArray = append(s.listArray, val)
	s.size++
}

func (s *AStack) Pop() int {
	top := s.size - 1
	ele := s.listArray[top]
	s.size--
	return ele
}

func (s AStack) TopValue() int {
	top := s.size - 1
	return s.listArray[top]
}

func (s AStack) Length() int {
	return s.size
}
