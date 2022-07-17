package utils

type List[T any] struct {
	Arr []T
	Cur int
}

func NewList[T any](size int) *List[T] {
	return &List[T]{make([]T, size), -1}
}

func (l *List[T]) Add(v T) {
	l.Cur++
	l.Arr[l.Cur] = v
}
