package searching

// contract: e's key < er's key
type node23 struct {
	// true=node3, false=node2
	is3    bool
	e      *entry
	er     *entry
	left   *node23
	mid    *node23
	right  *node23
	parent *node23

	l3cnt int
	r3cnt int
}

func (n *node23) isLeaf() bool {
	return n.left == nil //contract: left right must be nil at the same time
}

func downTo2(n *node23) {
	n.is3, n.er, n.mid = false, nil, nil
}

func (n *node23) getVal(k int) (v interface{}) {
	if n == nil {
		return nil
	}
	if n.is3 {
		if k < n.e.k {
			v = n.left.getVal(k)
		} else if k == n.e.k {
			v = n.e.v
		} else if k > n.e.k && k < n.er.k {
			v = n.mid.getVal(k)
		} else if k == n.er.k {
			v = n.er.v
		} else {
			v = n.right.getVal(k)
		}
	} else {
		if k < n.e.k {
			v = n.left.getVal(k)
		} else if k == n.e.k {
			v = n.e.v
		} else {
			v = n.right.getVal(k)
		}
	}
	return
}

func (n *node23) putVal(k int, v interface{}) (old interface{}) {
	if n.is3 {
		if k == n.e.k {
			old, n.e.v = n.e.v, v
		} else if k == n.er.k {
			old, n.er.v = n.er.v, v
		} else if n.isLeaf() {
			e := &entry{k, v}
			ascendMidToParentFromNode3(n, nil, e)
		} else {
			if k < n.e.k {
				old = n.left.putVal(k, v)
			} else if k > n.e.k && k < n.er.k {
				old = n.mid.putVal(k, v)
			} else {
				old = n.right.putVal(k, v)
			}
		}
	} else {
		if k == n.e.k {
			old, n.e.v = n.e.v, v
		} else if n.isLeaf() {
			upgradeLeafNode2(n, &entry{k, v})
		} else {
			if k < n.e.k {
				old = n.left.putVal(k, v)
			} else {
				old = n.right.putVal(k, v)
			}
		}
	}
	return
}

func (n *node23) removeVal(k int) (old interface{}) {
	if n.isLeaf() {
		old = removeFromLeaf(n, k)
	} else {
		if n.is3 {
			if n.e.k == k {
				s := swapInOrderSuccessor(n, LEFT)
				old = removeFromLeaf(s, k)
			} else if n.er.k == k {
				s := swapInOrderSuccessor(n, RIGHT)
				old = removeFromLeaf(s, k)
			} else if k < n.e.k {
				old = n.left.removeVal(k)
			} else if k > n.er.k {
				old = n.right.removeVal(k)
			} else {
				old = n.mid.removeVal(k)
			}
		} else {
			if n.e.k == k {
				s := swapInOrderSuccessor(n, 0)
				old = removeFromLeaf(s, k)
			} else if k < n.e.k {
				old = n.left.removeVal(k)
			} else {
				old = n.right.removeVal(k)
			}
		}
	}
	return
}

/* ----- utils ------ */

// find smallest key entry and its node
func floorNode23Tree(t *node23) *node23 {
	if t.isLeaf() {
		return t
	} else {
		return floorNode23Tree(t.left)
	}
}

// put n to a's position, n is empty node
func replaceNode23(n, a *node23) {
	n.e, n.er, n.is3 = a.e, a.er, a.is3
	connect(n, a.left, LEFT)
	connect(n, a.mid, MID)
	connect(n, a.right, RIGHT)
}
