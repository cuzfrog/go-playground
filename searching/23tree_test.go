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
	left, mid, right, cr := &node23{}, &node23{}, &node23{}, &node23{}

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
		assert.Equal(t, 5, nr2.e.k)
	})

	t.Run("right entry", func(t *testing.T) {
		n3 := aTestNode3()
		r := &entry{8, "right entry"}
		nr3, eu3 := splitNode3(n3, cr, r)
		assert.Equal(t, 2, n3.e.k)
		assert.Equal(t, 8, nr3.e.k)
		assert.Equal(t, 5, eu3.k)
		assert.Equal(t, cr, nr3.right)
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
	cl, cr := &node23{}, &node23{}
	n.left, n.right = cl, cr
	cl.parent, cr.parent = n, n
	nr := &node23{}
	eu := &entry{3, "as"}
	liftNode2ToRoot(n, nr, eu)
	assert.Equal(t, n, nr.parent)
	assert.Equal(t, nr, n.right)
	assert.Equal(t, n, n.left.parent)
	assert.Equal(t, 3, n.e.k)
	assert.Equal(t, 2, n.left.e.k)
	assert.Equal(t, cl, n.left.left)
	assert.Equal(t, cr, n.left.right)
	assert.Equal(t, n.left, cl.parent)
	assert.Equal(t, n.left, cr.parent)
}

func Test_ascendMidToParentFromNode3(t *testing.T) {
	t.Run("single root node3", func(t *testing.T) {
		n := aTestNode3()
		ascendMidToParentFromNode3(n, nil, &entry{8, "hes"})
		assert.False(t, n.is3)
		assert.Equal(t, 5, n.e.k)
		assert.Equal(t, 2, n.left.e.k)
		assert.Equal(t, 8, n.right.e.k)
		assert.Equal(t, n, n.right.parent)
	})

	left, mid, right, cr := &node23{}, &node23{}, &node23{}, &node23{}

	t.Run("node2 as parent", func(t *testing.T) {
		p := newNode2(8, "asgf", nil)
		n := aTestNode3()
		n.parent, p.left = p, n
		n.left, n.mid, n.right = left, mid, right
		ascendMidToParentFromNode3(n, cr, &entry{4, "gh"})
		assert.True(t, p.is3)
		assert.Equal(t, 4, p.e.k)
		assert.Equal(t, 8, p.er.k)
		assert.Equal(t, n, p.left)
		assert.False(t, n.is3)
		assert.Equal(t, 2, p.left.e.k)
		assert.Equal(t, 5, p.mid.e.k)
		assert.Equal(t, p, n.parent)
		assert.Equal(t, p, p.mid.parent)
		assert.Equal(t, left, p.left.left)
		assert.Equal(t, mid, p.left.right)
		assert.Equal(t, cr, p.mid.left)
		assert.Equal(t, right, p.mid.right)
	})

	t.Run("node3 as parent and root", func(t *testing.T) {
		p := newNode3(-3, "asgf", 1, "yd", nil)
		p.left, p.mid = &node23{parent: p}, &node23{parent: p}
		n := aTestNode3()
		n.parent, p.right = p, n
		n.left, n.mid, n.right = left, mid, right
		ascendMidToParentFromNode3(n, cr, &entry{4, "ghii"})
		assert.False(t, p.is3)
		assert.False(t, n.is3)
		assert.Equal(t, n, p.right.left)
		assert.Equal(t, 1, p.e.k)
		assert.Equal(t, -3, p.left.e.k)
		assert.Equal(t, 4, p.right.e.k)
		assert.Equal(t, 2, n.e.k)
		assert.Equal(t, 5, p.right.right.e.k)
		assert.Equal(t, p.right, n.parent)
	})
}

func Test_borrowDownward(t *testing.T) {
	asert := assert.New(t)
	assertConnect := func(p, c *node23, pos position) {
		if pos == LEFT {
			asert.Equal(p.left, c)
			asert.Equal(p, c.parent)
		} else if pos == MID {
			asert.Equal(p.mid, c)
			asert.Equal(p, c.parent)
		} else {
			asert.Equal(p.right, c)
			asert.Equal(p, c.parent)
		}
	}
	/*
		 2|5            3
	   /  |  \        /  \
	  x   3  7       2  5|7
	 /   /\  /\     /\  /|\
	a   d f g i    a d f g i
	*/
	t.Run("node3 parent / node2 siblings L", func(t *testing.T) {
		p := aTestTree()
		downTo2(p.left)
		downTo2(p.mid)
		downTo2(p.right)
		borrowDownward(p.left)
		asert.False(p.is3)
		asert.Equal(3, p.e.k)
		assertConnect(p, p.left, LEFT)
		asert.Nil(p.mid)
		assertConnect(p, p.right, RIGHT)
		asert.Equal(2, p.left.e.k)
		asert.Equal(5, p.right.e.k)
		asert.Equal(7, p.right.er.k)
		asert.Equal("a", p.left.left.e.v)
		asert.Equal("d", p.left.right.e.v)
		asert.Equal("f", p.right.left.e.v)
		asert.Equal("g", p.right.mid.e.v)
		asert.Equal("i", p.right.right.e.v)
	})

	/*
		 2|5
	   /  |  \
	  0   3  7
	 /\  /\  /\
	a c d f g i
	*/
	t.Run("node3 parent / node2 siblings M", func(t *testing.T) {
		p := aTestTree()
		downTo2(p.left)
		downTo2(p.mid)
		downTo2(p.right)
		borrowDownward(p.left)
	})

	t.Run("node3 parent / node3 neighbor sibling", func(t *testing.T) {

	})

	t.Run("node3 parent / node3 remote sibling", func(t *testing.T) {

	})

	t.Run("node2 parent / node3 sibling", func(t *testing.T) {

	})

	t.Run("node2 parent / node2 sibling", func(t *testing.T) {

	})
}

// {2,5}
func aTestNode3() *node23 {
	return newNode3(2, utils.RandAlphabet(2), 5, utils.RandAlphabet(3), nil)
}

// {2}
func aTestNode2() *node23 {
	return newNode2(2, utils.RandAlphabet(2), nil)
}

func newNode2(k int, v interface{}, p *node23) *node23 {
	return &node23{e: &entry{k, v}, parent: p}
}

func newNode3(k1 int, v1 interface{}, k2 int, v2 interface{}, p *node23) *node23 {
	return &node23{is3: true, e: &entry{k1, v1}, er: &entry{k2, v2}, parent: p}
}

/*
     2|5
   /  |  \
 0|1 3|4 7|8
 /|\ /|\ /|\
 abc def ghi
*/
func aTestTree() *node23 {
	p := aTestNode3()
	l := newNode3(0, "l1", 1, "l2", nil)
	m := newNode3(3, "m1", 4, "m2", nil)
	r := newNode3(7, "r1", 8, "r2", nil)
	connect(p, l, LEFT)
	connect(p, m, MID)
	connect(p, r, RIGHT)
	a := &node23{e: &entry{v: "a"}}
	b := &node23{e: &entry{v: "b"}}
	c := &node23{e: &entry{v: "c"}}
	d := &node23{e: &entry{v: "d"}}
	e := &node23{e: &entry{v: "e"}}
	f := &node23{e: &entry{v: "f"}}
	g := &node23{e: &entry{v: "g"}}
	h := &node23{e: &entry{v: "h"}}
	i := &node23{e: &entry{v: "i"}}
	connect(l, a, LEFT)
	connect(l, b, MID)
	connect(l, c, RIGHT)
	connect(m, d, LEFT)
	connect(m, e, MID)
	connect(m, f, RIGHT)
	connect(r, g, LEFT)
	connect(r, h, MID)
	connect(r, i, RIGHT)
	return p
}
