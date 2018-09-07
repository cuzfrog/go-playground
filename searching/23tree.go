package searching

import "fmt"

type twoThreeTree struct {
	count int
	root  nodeP
}

func (t *twoThreeTree) get(k int) interface{} {
	return t.root.getVal(k)
}

func (t *twoThreeTree) put(k int, v interface{}) (old interface{}, err error) {
	if v == nil {
		return nil, fmt.Errorf("v cannot be nil")
	}
	if t.root == nil {
		t.root = &node2{e: entry{k, v}}
	} else {
		old, t.root = t.root.putVal(k, v)
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

func upgradeNode2(n *node2, e *entry) *node3 {
	if e.k < n.e.k {
		return &node3{el: *e, er: n.e, parent: n.parent}
	} else {
		return &node3{el: n.e, er: *e, parent: n.parent}
	}
}

func splitRootNode3(n *node3, e *entry) (mid *node2) {
	var left, right *node2
	if e.k < n.el.k {
		left = &node2{e: *e}
		right = &node2{e: n.er}
		mid = &node2{e: n.el}
	} else if e.k > n.el.k && e.k < n.er.k {
		left = &node2{e: n.el}
		right = &node2{e: n.er}
		mid = &node2{e: *e}
	} else if e.k > n.el.k {
		left = &node2{e: n.el}
		right = &node2{e: *e}
		mid = &node2{e: n.er}
	} else {
		panic("duplicate key when splitting root node3")
	}
	left.parent, right.parent = mid, mid
	mid.left, mid.right = left, right
	return
}
