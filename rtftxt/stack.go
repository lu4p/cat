package rtftxt

type stack struct {
	top  *element
	size int
}

type element struct {
	value string
	next  *element
}

func (s *stack) Len() int {
	return s.size
}

func (s *stack) Push(value string) {
	s.top = &element{value, s.top}
	s.size++
}

func (s *stack) Peek() string {
	if s.size == 0 {
		return ""
	}
	return s.top.value
}

func (s *stack) Pop() string {
	if s.size > 0 {
		var v string
		v, s.top = s.top.value, s.top.next
		s.size--
		return v
	}
	return ""
}
