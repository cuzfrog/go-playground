package searching

import "fmt"

type twoThreeTree struct {
	count int
	root  *node23
}

func (t *twoThreeTree) get(k int) interface{} {
	return t.root.getVal(k)
}

func (t *twoThreeTree) put(k int, v interface{}) (old interface{}, err error) {
	if v == nil {
		return nil, fmt.Errorf("v cannot be nil")
	}
	if t.root == nil {
		t.root = &node23{e: &entry{k, v}}
	} else {
		old = t.root.putVal(k, v)
	}
	t.count++
	return
}

func (*twoThreeTree) remove(k int) interface{} {
	panic("implement me")
}

func (t *twoThreeTree) contains(k int) bool {
	return t.get(k) != nil
}

func (t *twoThreeTree) size() int {
	return t.count
}

func (*twoThreeTree) iterator() chan entry {
	panic("implement me")
}

/* -------------- 23tree transformation implementation -------------- */

func upgradeLeafNode2(n *node23, e *entry) {
	if e.k < n.e.k {
		n.er, n.e = n.e, e
	} else if e.k > n.e.k {
		n.er = e
	} else {
		panic("duplicate key when upgrading node2")
	}
	n.is3 = true
}

func splitRootLeafNode3(n *node23, e *entry) {
	var left, right *node23
	if e.k < n.e.k {
		left = &node23{e: e}
		right = &node23{e: n.er}
	} else if e.k > n.e.k && e.k < n.er.k {
		left = &node23{e: n.e}
		right = &node23{e: n.er}
		n.e = e
	} else if e.k > n.e.k {
		left = &node23{e: n.e}
		right = &node23{e: e}
		n.e = n.er
	} else {
		panic("duplicate key when splitting root node23")
	}
	left.parent, right.parent = n, n
	n.left, n.right = left, right
	n.is3, n.er = false, nil
}

func ascendMidToParentFromNode3(n *node23, e *entry) {
	var eu *entry
	if e.k < n.e.k {
		eu, n.e = n.e, e
	} else if e.k > n.e.k && e.k < n.er.k {
		eu = e
	} else if e.k > n.e.k {
		eu, n.er = n.er, e
	} else {
		panic("duplicate key when ascending mid to parent node from leaf node23")
	}
	nl, nr := n, &node23{e: n.er}
	n.is3 = false

	p := n.parent
	if p != nil {
		if p.is3 {
			ascendMidToParentFromNode3(p, eu)
		} else {
			if p.left == n {
				p.e, p.er, p.left, p.mid = eu, p.e, nl, nr
			} else if p.right == n {
				p.er, p.mid, p.right = eu, nl, nr
			} else {
				panic(fmt.Sprintf("node[%v] is not a child of its parent node[%v]", *n, *p))
			}
			nl.parent, nr.parent = p, p
			p.is3 = true
		}
	} else { //root

	}
}
