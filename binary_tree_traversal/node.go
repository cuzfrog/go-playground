package binary_tree_traversal

const left, right = true, false

type node struct {
	v int
	a *node
	b *node
}

func newNode(v int, p *node, branch bool) *node {
	n := &node{v, nil, nil}
	if p != nil {
		if branch == left {
			p.a = n
		} else {
			p.b = n
		}
	}
	return n
}
