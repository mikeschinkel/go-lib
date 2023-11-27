package blank_identifier_alias_this_package

type Stack[T comparable] struct {
	stack []T
}

func (s *Stack[T]) Push(v T) {
	s.stack = append(s.stack, v)
}

func (s *Stack[T]) Has(v T) (has bool) {
	for _, e := range s.stack {
		if e == v {
			return true
		}
	}
	return false
}

func (s *Stack[T]) Pop() T {
	return s.stack[len(s.stack)-1]
}
func (s *Stack[T]) Drop() {
	s.stack = s.stack[:len(s.stack)-1]
}
func (s *Stack[T]) Empty() bool {
	return len(s.stack) == 0
}
func (s *Stack[T]) Depth() int {
	return len(s.stack)
}

func (s *Stack[T]) Top() T {
	return s.stack[len(s.stack)-1]
}
