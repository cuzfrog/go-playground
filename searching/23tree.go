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

// splitNode3 split a node3 to two node2s, the left of which is the original node
//	parameters:
//	 n - the current node3
//	 cr - the newly split right sibling of one child, nil if n is leaf
//   e - mid entry sent by one of the children node3s or new entry inserted by client
//  return:
//	 nr - newly created right sibling node2 of n
//   eu - calculated mid entry
func splitNode3(n, cr *node23, e *entry) (nr *node23, eu *entry) {
	nr = &node23{parent: n.parent}
	if e.k < n.e.k {
		nr.e = n.er
		eu, n.e = n.e, e
		nr.left, nr.right = n.mid, n.right
		n.right = cr
	} else if e.k > n.e.k && e.k < n.er.k {
		eu = e
		nr.left, nr.right = cr, n.right
		n.right = n.mid
	} else if e.k > n.e.k {
		eu, n.er = n.er, e
		nr.left, nr.right = n.right, cr
		n.right = n.mid
	} else {
		panic("duplicate key when splitting node3")
	}
	downTo2(n)
	return
}

// ascendMidToParentFromNode3 inserts mid entry into parent node from a node3
//  parameters:
//   n - current node3 from which a mid entry is calculated by its existing entries and a given entry 'e'
//   cr - the newly split right sibling of one child, nil if n is leaf
//   e - mid entry sent by one of the children node3s or new entry inserted by client
func ascendMidToParentFromNode3(n, cr *node23, e *entry) {
	nr, eu := splitNode3(n, cr, e)
	p := n.parent
	if p != nil {
		if p.is3 {
			ascendMidToParentFromNode3(p, nr, eu)
		} else {
			insertMidFromChildToNode2(n, nr, eu)
		}
	} else { //n is root node3
		liftNode2ToRoot(n, nr, eu)
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

// liftNode2ToRoot lifts previous root to new root with mid entry
func liftNode2ToRoot(n, nr *node23, eu *entry) {
	nl := &node23{e: n.e, left: n.left, right: n.right, parent: n.parent}
	n.left.parent, n.right.parent = nl, nl
	n.left, n.right, n.e = nl, nr, eu
}
