package searching

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"go-playground/utils"
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

func Test_splitNode3(t *testing.T) {
	t.Run("left entry", func(t *testing.T) {
		n1 := aTestNode3()
		l := &entry{1, "left entry"}
		nr1, eu1 := splitNode3(n1, nil, l)
		assert.Equal(t, 1, n1.e.k)
		assert.False(t, n1.is3)
		assert.Nil(t, n1.mid)
		assert.Nil(t, n1.er)
		assert.Equal(t, 2, eu1.k)
		assert.Equal(t, 5, nr1.e.k)
	})

	left, mid, right, cr := &node23{}, &node23{}, &node23{}, &node23{}

	t.Run("mid entry", func(t *testing.T) {
		n2 := aTestNode3()
		n2.left, n2.mid, n2.right = left, mid, right
		m := &entry{3, "mid entry"}
		nr2, _ := splitNode3(n2, cr, m)
		assert.Equal(t, 2, n2.e.k)
		assert.Equal(t, left, n2.left)
		assert.Equal(t, mid, n2.right)
		assert.Equal(t, cr, nr2.left)
		assert.Equal(t, right, nr2.right)
	})

	t.Run("right entry", func(t *testing.T) {
		n3 := aTestNode3()
		r := &entry{8, "right entry"}
		splitNode3(n3, nil, r)
		assert.Equal(t, 2, n3.e.k)
	})
}

func Test_insertMidFromChildToNode2(t *testing.T) {
	n := aTestNode2()
	cr := &node23{}
	eu := &entry{4, "mid"}
	t.Run("left branch", func(t *testing.T) {
		p1 := newNode2(7, "pp", nil)
		n.parent, p1.left = p1, n
		insertMidFromChildToNode2(n, cr, eu)
		assert.True(t, p1.is3)
		assert.Equal(t, 4, p1.e.k)
		assert.Equal(t, 7, p1.er.k)
		assert.Equal(t, n, p1.left)
		assert.Equal(t, cr, p1.mid)
	})

	t.Run("right branch", func(t *testing.T) {
		p2 := newNode2(-1, "pp", nil)
		n.parent, p2.right = p2, n
		insertMidFromChildToNode2(n, cr, eu)
		assert.True(t, p2.is3)
		assert.Equal(t, -1, p2.e.k)
		assert.Equal(t, 4, p2.er.k)
		assert.Equal(t, n, p2.mid)
		assert.Equal(t, cr, p2.right)
	})
}

func Test_liftNode2ToRoot(t *testing.T) {
	n := aTestNode2()
	nr := &node23{}
	eu := &entry{3, "as"}
	liftNode2ToRoot(n, nr, eu)
	assert.Equal(t, n, nr.parent)
	assert.Equal(t, nr, n.right)
	assert.Equal(t, n, n.left.parent)
	assert.Equal(t, 3, n.e.k)
	assert.Equal(t, 2, n.left.e.k)
}

func aTestNode3() *node23 {
	return newNode3(2, utils.RandAlphabet(2), 5, utils.RandAlphabet(3), nil)
}

func aTestNode2() *node23 {
	return newNode2(2, utils.RandAlphabet(2), nil)
}

func newNode2(k int, v interface{}, p *node23) *node23 {
	return &node23{e: &entry{k, v}, parent: p}
}

func newNode3(k1 int, v1 interface{}, k2 int, v2 interface{}, p *node23) *node23 {
	return &node23{e: &entry{k1, v1}, er: &entry{k2, v2}, parent: p}
}
