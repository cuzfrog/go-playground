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

func (t *twoThreeTree) remove(k int) (old interface{}) {
	if t.root == nil {
		return nil
	}
	if t.root.isLeaf() && !t.root.is3 && t.root.e.k == k { //remove single node2
		old = t.root.e.v
		t.root = nil
	} else {
		old = t.root.removeVal(k)
	}
	if old != nil {
		t.count--
	}
	return
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

type position byte

const (
	LEFT  position = iota
	MID
	RIGHT
)

// connect establishes parent-child relationship
func connect(p, c *node23, pos position) {
	if pos == LEFT {
		p.left = c
	} else if pos == MID {
		p.mid = c
	} else {
		p.right = c
	}
	if c != nil {
		c.parent = p
	}
}

// disconnect removes parent-child relationship
func disconnect(p, c *node23, pos position) {
	if pos == LEFT {
		p.left = nil
	} else if pos == MID {
		p.mid = nil
	} else {
		p.right = nil
	}
	if c != nil {
		c.parent = nil
	}
}

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
		connect(nr, n.mid, LEFT)
		connect(nr, n.right, RIGHT)
		connect(n, cr, RIGHT)
	} else if e.k > n.e.k && e.k < n.er.k {
		nr.e = n.er
		eu = e
		connect(nr, cr, LEFT)
		connect(nr, n.right, RIGHT)
		connect(n, n.mid, RIGHT)
	} else if e.k > n.e.k {
		nr.e = e
		eu, n.er = n.er, e
		connect(nr, n.right, LEFT)
		connect(nr, cr, RIGHT)
		connect(n, n.mid, RIGHT)
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
		p.e, p.er = eu, p.e
		connect(p, nl, LEFT)
		connect(p, cr, MID)
	} else if p.right == n {
		p.er = eu
		connect(p, nl, MID)
		connect(p, cr, RIGHT)
	} else {
		panic(fmt.Sprintf("node[%v] is not a child of its parent node[%v]", *n, *p))
	}
	p.is3 = true
}

// liftNode2ToRoot lifts previous root to new root with mid entry
func liftNode2ToRoot(n, nr *node23, eu *entry) {
	nl := &node23{e: n.e, left: n.left, right: n.right, parent: n}
	if !n.isLeaf() {
		n.left.parent, n.right.parent = nl, nl
	}
	n.left, n.right, n.e = nl, nr, eu
	nr.parent = n
}

//todo: test
func swapInOrderSuccessor(n *node23, pos position) (s *node23) {
	if n.isLeaf() {
		panic("contract violated")
	}
	if n.is3 {
		if pos == LEFT {
			s = floor(n.mid)
			s.e, n.e = n.e, s.e
		} else if pos == RIGHT {
			s = floor(n.right)
			s.e, n.er = n.er, s.e
		} else {
			panic("pos stands for entry position")
		}
	} else {
		s = floor(n.right)
		s.e, n.e = n.e, s.e
	}
	return
}

func removeFromLeaf(n *node23, k int) (old interface{}) {
	if n.is3 {
		old = removeFromLeafNode3(n, k)
	} else {
		old = removeFromLeafNode2(n, k)
	}
	return
}

//todo: test
func removeFromLeafNode3(n *node23, k int) (old interface{}) {
	if k == n.e.k {
		old = n.e.v
		n.e = n.er
		downTo2(n)
	} else if k == n.er.k {
		old = n.er.v
		downTo2(n)
	}
	return
}

//todo: test
func removeFromLeafNode2(n *node23, k int) (old interface{}) {
	if n.e.k != k {
		return
	}
	old = n.e.v
	borrowDownward(n)
	return
}

//todo: test
func borrowDownward(h *node23) {
	p := h.parent
	if p == nil {

	} else if p.is3 {
		if p.left == h {
			sm, sr := p.mid, p.right
			if !sm.is3 && !sr.is3 { //node2 siblings
				h.e, p.e, sr.er, sr.e = p.e, sm.e, sr.e, p.er
				connect(h, sm.left, RIGHT)
				connect(sr, sr.left, MID)
				connect(sr, sm.right, LEFT)
				sr.is3 = true
				downTo2(p)
			} else if sm.is3 {
				h.e, p.e, sm.e = p.e, sm.e, sm.er
				connect(h, sm.left, RIGHT)
				connect(sm, sm.mid, LEFT)
				downTo2(sm)
			} else { //only sr is node3
				h.e, p.e, sm.e, p.er, sr.e = p.e, sm.e, p.er, sr.e, sr.er
				connect(h, sm.left, RIGHT)
				connect(sm, sm.right, LEFT)
				connect(sm, sr.left, RIGHT)
				connect(sr, sr.mid, LEFT)
				downTo2(sr)
			}
		} else if p.mid == h {
			sl, sr := p.left, p.right
			if !sl.is3 && !sr.is3 { //node2 siblings
				sl.er, p.e = p.e, p.er
				connect(sl, sl.right, MID)
				connect(sl, h.left, RIGHT)
				sl.is3 = true
				downTo2(p)
			} else if sl.is3 {
				h.e, p.e = p.e, sl.er
				connect(h, h.left, RIGHT)
				connect(h, sl.right, LEFT)
				connect(sl, sl.mid, RIGHT)
				downTo2(sl)
			} else { //only sr is node3
				h.e, p.er, sr.e = p.er, sr.e, sr.er
				connect(h, sr.left, RIGHT)
				connect(sr, sr.mid, LEFT)
				downTo2(sr)
			}
		} else {
			sl, sm := p.left, p.mid
			if !sl.is3 && !sm.is3 { //node2 siblings
				sl.er, p.e, h.e = p.e, sm.e, p.er
				connect(sl, sl.right, MID)
				connect(sl, sm.left, RIGHT)
				sl.is3 = true
				connect(h, h.left, RIGHT)
				connect(h, sm.right, LEFT)
				downTo2(p)
			} else if sm.is3 {
				h.e, p.er = p.er, sm.er
				connect(h, h.left, RIGHT)
				connect(h, sm.right, LEFT)
				connect(sm, sm.mid, RIGHT)
				downTo2(sm)
			} else { //only sl is node3
				h.e, p.er, sm.e, p.e = p.er, sm.e, p.e, sl.er
				connect(h, h.left, RIGHT)
				connect(h, sm.right, LEFT)
				connect(sm, sm.left, RIGHT)
				connect(sm, sl.right, LEFT)
				connect(sl, sl.mid, RIGHT)
				downTo2(sl)
			}
		}
	} else { //parent is node2
		if p.left == h {
			s := p.right
			if s.is3 {
				h.e, p.e, s.e = p.e, s.e, s.er
				connect(h, s.left, RIGHT)
				connect(s, s.mid, LEFT)
				downTo2(s)
			} else {
				h.e = p.e
				h.er = s.e
				connect(h, s.left, MID)
				connect(h, s.right, RIGHT)
				h.is3 = true
				disconnect(p, s, RIGHT)
				borrowDownward(p)
			}
		} else {
			s := p.left
			if s.is3 {
				h.e, p.e = p.e, s.er
				connect(h, h.left, RIGHT)
				connect(h, s.right, LEFT)
				connect(s, s.mid, RIGHT)
				downTo2(s)
			} else {
				s.er = p.e
				connect(s, s.right, MID)
				connect(s, h.left, RIGHT)
				s.is3 = true
				disconnect(p, h, RIGHT)
				borrowDownward(p)
			}
		}
	}
}
