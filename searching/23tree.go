package searching

import "fmt"

type twoThreeTree struct {
	count int
	root  nodeP
}

func (t *twoThreeTree) get(k int) interface{} {
	return t.root.getVal(k)
}

func (t *twoThreeTree) put(k int, v interface{}) (interface{}, error) {
	if v == nil {
		return nil, fmt.Errorf("v cannot be nil")
	}
	if t.root == nil {
		t.root = &node2{e: entry{k, v}}
	}
	t.count++

	switch r := t.root.(type) {
	
	}
	return t.root.putVal(k, v)
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
		return &node3{el: *e, er: n.e}
	} else {
		return &node3{el: n.e, er: *e}
	}
}
