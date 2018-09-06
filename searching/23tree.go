package searching

type twoThreeTree struct {
	count int
	root  *node23
}

type node23 interface{}

type node2 struct {
	e     entry
	left  *node23
	right *node23
}

// contract: el's key < er's key
type node3 struct {
	el    entry
	er    entry
	left  *node23
	mid   *node23
	right *node23
}

func (t *twoThreeTree) get(k int) interface{} {
	return get(t.root, k)
}

func get(np *node23, k int) (v interface{}) {
	if np == nil {
		return nil
	}
	switch n := (*np).(type) {
	case node2:
		if n.e.k == k {
			v = n.e.v
		} else if k < n.e.k {
			v = get(n.left, k)
		} else {
			v = get(n.right, k)
		}
	case node3:
		if k < n.el.k {
			v = get(n.left, k)
		} else if k == n.el.k {
			v = n.el.v
		} else if k > n.el.k && k < n.er.k {
			v = get(n.mid, k)
		} else if k == n.er.k {
			v = n.er.v
		} else {
			v = get(n.right, k)
		}
	}
	return
}

func (*twoThreeTree) put(k int, v interface{}) {
	panic("implement me")
}

func (*twoThreeTree) remove(k int) interface{} {
	panic("implement me")
}

func (*twoThreeTree) contains(k int) bool {
	panic("implement me")
}

func (t *twoThreeTree) size() int {
	return t.count
}

func (*twoThreeTree) iterator() chan entry {
	panic("implement me")
}
