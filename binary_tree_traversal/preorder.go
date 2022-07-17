package binary_tree_traversal

import "github.com/cuzfrog/go-playground/utils"

func recursivePreorderTraverse(n *node, l *utils.List[int]) {
	l.Add(n.v)
	s := utils.NewStack[*node](len(l.Arr))
	if n.b != nil {
		s.Push(n.b)
	}
	for n.a != nil {
		n = n.a
		l.Add(n.v)
		if n.b != nil {
			s.Push(n.b)
		}
	}
	for s.Size() > 0 {
		b := s.Pop()
		recursivePreorderTraverse(b, l)
	}
}
