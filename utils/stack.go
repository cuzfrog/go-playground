package utils

type Stack[T any] struct {
	Arr []T
	Cur int
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{make([]T, size), -1}
}

func (s *Stack[T]) Push(v T) {
	s.Cur++
	s.Arr[s.Cur] = v
}

func (s *Stack[T]) Pop() T {
	v := s.Arr[s.Cur]
	s.Cur--
	return v
}

func (s *Stack[T]) Size() int {
	return s.Cur + 1
}
