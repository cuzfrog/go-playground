package searching

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_upgradeLeafNode2(t *testing.T) {
	p := &node23{}
	n1 := newNode2(1, "v1", p)
	upgradeLeafNode2(n1, &entry{2, "right"})
	assert.Equal(t, "right", n1.er.v)
	assert.Equal(t, p, n1.parent)
	assert.True(t, n1.is3)

	n2 := newNode2(1, "v2", p)
	upgradeLeafNode2(n2, &entry{-1, "left"})
	assert.Equal(t, 1, n2.er.k)
	assert.Equal(t, -1, n2.e.k)
}

func Test_splitRootLeafNode3(t *testing.T) {
	n1 := newNode3(2, "s", 5, "gg", nil)
	l := &entry{1, "left entry"}
	splitRootLeafNode3(n1, l)
	assert.Equal(t, 2, n1.e.k)
	assert.False(t, n1.is3)
	assert.Equal(t, 1, n1.left.e.k)
	assert.Equal(t, 5, n1.right.e.k)

	n2 := newNode3(2, "s2", 5, "gg3", nil)
	m := &entry{3, "mid entry"}
	splitRootLeafNode3(n2, m)
	assert.Equal(t, 3, n2.e.k)

	n3 := newNode3(2, "s3", 5, "gg3g", nil)
	r := &entry{8, "right entry"}
	splitRootLeafNode3(n3, r)
	assert.Equal(t, 5, n3.e.k)
}

func newNode2(k int, v interface{}, p *node23) *node23 {
	return &node23{e: &entry{k, v}, parent: p}
}

func newNode3(k1 int, v1 interface{}, k2 int, v2 interface{}, p *node23) *node23 {
	return &node23{e: &entry{k1, v1}, er: &entry{k2, v2}, parent: p}
}
