package stack

type Link struct {
	element int
	next    *Link
}

type LStack struct {
	size int
	top  *Link
}

func (s *LStack) Clear() {
	for s.top != nil {
		s.top = s.top.next
	}
	s.size = 0
}

func (s *LStack) Push(val int) {
	s.top = &Link{element: val, next: s.top}
	s.size++
}

func (s *LStack) Pop() int {
	val := s.top.element
	s.top = s.top.next
	s.size--
	return val
}

func (s LStack) TopValue() int {
	return s.top.element
}

func (s LStack) Length() int {
	return s.size
}


