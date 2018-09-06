package searching

import "fmt"

type nodeP interface {
	getVal(k int) interface{}
	putVal(k int, v interface{}) (interface{}, error)
}

type node2 struct {
	e     entry
	left  nodeP
	right nodeP
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

func (n *node2) putVal(k int, v interface{}) (old interface{}, err error) {
	if n == nil {
		panic("putVal is called on nil node")
	}
	if v == nil {
		return nil, fmt.Errorf("v cannot be nil")
	}
	if k < n.e.k {
		old, err = n.left.putVal(k, v)
	} else if k == n.e.k {
		old = n.e.v
		n.e.v = v
	} else {
		old, err = n.right.putVal(k, v)
	}
	return
}

// contract: el's key < er's key
type node3 struct {
	el    entry
	er    entry
	left  nodeP
	mid   nodeP
	right nodeP
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
