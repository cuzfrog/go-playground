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

// upgradeLeafNode2 inserts a new entry 'e' into node2 'n' and turns it into a node3
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

// splitRootNode3 split a node3 at root
//	parameters:
//	 n - the root node3
//	 cr - the newly split right sibling of one child, nil if n is leaf
//   e - mid entry sent by one of the children node3s or new entry inserted by client
func splitRootNode3(n, cr *node23, e *entry) {
	var left, right *node23
	if e.k < n.e.k {
		left = &node23{e: e}
		right = &node23{e: n.er}
		left.left, left.right = n.left, cr
		right.left, right.right = n.mid, n.right
	} else if e.k > n.e.k && e.k < n.er.k {
		left = &node23{e: n.e}
		right = &node23{e: n.er}
		n.e = e
		left.left, left.right = n.left, n.mid
		right.left, right.right = cr, n.right
	} else if e.k > n.e.k {
		left = &node23{e: n.e}
		right = &node23{e: e}
		n.e = n.er
		left.left, left.right = n.left, n.mid
		right.left, right.right = n.right, cr
	} else {
		panic("duplicate key when splitting root node23")
	}
	left.parent, right.parent = n, n
	n.left, n.right = left, right
	downTo2(n)
}

// ascendMidToParentFromNode3 inserts mid entry into parent node from a node3
//  parameters:
//   n - current node3 from which a mid entry is calculated by its existing entries and a given entry 'e'
//   cr - the newly split right sibling of one child, nil if n is leaf
//   e - mid entry sent by one of the children node3s or new entry inserted by client
func ascendMidToParentFromNode3(n, cr *node23, e *entry) {
	var eu *entry //mid entry
	if e.k < n.e.k {
		eu, n.e = n.e, e
	} else if e.k > n.e.k && e.k < n.er.k {
		eu = e
	} else if e.k > n.e.k {
		eu, n.er = n.er, e
	} else {
		panic("duplicate key when ascending mid to parent node from leaf node23")
	}
	nl, nr := n, &node23{e: n.er} //split right sibling
	nr.parent, nr.left, nr.right = n.parent, n.mid, n.right
	nl.right = cr
	downTo2(n)

	p := n.parent
	if p != nil {
		if p.is3 {
			ascendMidToParentFromNode3(p, nr, eu)
		} else {
			insertMidFromChildToNode2(n, nr, eu)
		}
	} else { //n is root node3
		splitRootNode3(n, nr, eu)
	}
}

// insertMidFromChildToNode2 inserts calculated mid entry to parent node2
//  parameters:
//   n - current child node3 which has already been downgraded to node2
//   cr - n's right sibling node2
//   eu - calculated mid entry to insert into parent node2
func insertMidFromChildToNode2(n, cr *node23, eu *entry) {
	nl, p := n, n.parent
	if p.left == n {
		p.e, p.er, p.left, p.mid = eu, p.e, nl, cr
	} else if p.right == n {
		p.er, p.mid, p.right = eu, nl, cr
	} else {
		panic(fmt.Sprintf("node[%v] is not a child of its parent node[%v]", *n, *p))
	}
	p.is3 = true
}
