package searching

type nodeP interface {
	getVal(k int) interface{}
	putVal(k int, v interface{}) (interface{}, nodeP)
	isLeaf() bool
}

type node2 struct {
	e      entry
	left   nodeP
	right  nodeP
	parent nodeP
}

func (n *node2) isLeaf() bool {
	return n.left == nil
}

func (n *node2) getVal(k int) (v interface{}) {
	if n == nil {
		return nil
	}
	if k < n.e.k {
		v = n.left.getVal(k)
	} else if k == n.e.k {
		v = n.e.v
	} else {
		v = n.right.getVal(k)
	}
	return
}

func (n *node2) putVal(k int, v interface{}) (old interface{}, u nodeP) {
	u = n
	if k == n.e.k {
		old, n.e.v = n.e.v, v
	} else if n.isLeaf() { //contract: left right must be nil at the same time
		u = upgradeNode2(n, &entry{k, v})
	} else {
		if k < n.e.k {
			old, n.left = n.left.putVal(k, v)
		} else {
			old, n.right = n.right.putVal(k, v)
		}
	}
	return
}

// contract: el's key < er's key
type node3 struct {
	el     entry
	er     entry
	left   nodeP
	mid    nodeP
	right  nodeP
	parent nodeP
}

func newNode3() {

}

func (n *node3) isLeaf() bool {
	return n.left == nil
}

func (n *node3) getVal(k int) (v interface{}) {
	if n == nil {
		return nil
	}
	if k < n.el.k {
		v = n.left.getVal(k)
	} else if k == n.el.k {
		v = n.el.v
	} else if k > n.el.k && k < n.er.k {
		v = n.mid.getVal(k)
	} else if k == n.er.k {
		v = n.er.v
	} else {
		v = n.right.getVal(k)
	}
	return
}

func (n *node3) putVal(k int, v interface{}) (old interface{}, u nodeP) {
	if k == n.el.k {
		old, n.el.v = n.el.v, v
	} else if k == n.er.k {
		old, n.er.v = n.er.v, v
	} else if n.isLeaf() {
		if n.parent == nil { //is root

		} else {

		}
	} else {
		if k < n.el.k {
			old, n.left = n.left.putVal(k, v)
		} else if k > n.el.k && k < n.er.k {
			old, n.mid = n.mid.putVal(k, v)
		} else {
			old, n.right = n.right.putVal(k, v)
		}
	}
	return
}
