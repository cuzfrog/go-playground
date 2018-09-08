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
