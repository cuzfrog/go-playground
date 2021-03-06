package searching

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_rbnode_functional(t *testing.T) {
	asert := assert.New(t)
	n := &rbnode{k: 10, v: "10"}
	nn, old := n.insert(5, "5")
	asert.Equal(nn, n)
	asert.Nil(old)
	asert.Equal(5, n.left.k)
	asert.Equal(red, n.left.c)
	asert.Equal(black, n.c)
	/*   10b
	     /
	    5r
	 */
	_, old = n.insert(5, "5a")
	asert.Equal("5", old)
	asert.Equal("5a", n.left.v)
	_, old = n.insert(10, "10a")
	asert.Equal("10", old)
	asert.Equal("10a", n.v)

	n.insert(11, "11")
	asert.Equal(black, n.left.c)
	asert.Equal(black, n.right.c)
	asert.Equal(red, n.c)
	asert.Equal(11, n.right.k)
	/*   10r
	    /  \
	   5b  11b
	 */
	_, old = n.insert(11, "11a")
	asert.Equal("11", old)
	asert.Equal("11a", n.right.v)

	n.insert(8, "8")
	asert.Equal(8, n.left.k)
	asert.Equal(5, n.left.left.k)
	asert.Equal(red, n.left.left.c)
	/*   10r         10r
	    /  \        /   \
	   5b  11b     8b   11b
	    \         /
	     8r     5r
	 */

	n.insert(6, "6")
	asert.Equal(red, n.left.c)
	asert.Equal(black, n.left.left.c)
	asert.Equal(black, n.left.right.c)
	asert.Equal(6, n.left.k)
	asert.Equal(5, n.left.left.k)
	asert.Equal(8, n.left.right.k)
	/*  10r         10r            10r
	   /  \        /   \          /   \
	  8b  11b     8b   11b       6r   11b
	 /           /             /   \
	5r          6r            5b   8b
	 \         /
	  6r      5r
	*/

	n.insert(2, "2")
	n, _ = n.insert(1, "1")
	asert.Equal(6, n.k)
	asert.Equal(red, n.c)
	asert.Equal(2, n.left.k)
	asert.Equal(black, n.left.c)
	asert.Equal(1, n.left.left.k)
	asert.Equal(black, n.left.left.c)
	asert.Equal(10, n.right.k)
	asert.Equal(black, n.right.c)
	asert.Equal(8, n.right.left.k)
	asert.Equal(black, n.right.left.c)
	asert.Equal(11, n.right.right.k)
	asert.Equal(black, n.right.right.c)
	/*    10r          6r
	     /  \        /   \
	    6r  11b     2b   10b
	   /  \        / \   / \
	  2r   8b     1b 5b 8b 11b
	 / \
	1b 5b
	 */

	_, old = n.remove(99)
	asert.Nil(old)

	n, old = n.remove(8)
	asert.Equal("8", old)
	asert.Equal(6, n.k)
	asert.Equal(black, n.c)
	asert.Equal(2, n.left.k)
	asert.Equal(red, n.left.c)
	asert.Equal(11, n.right.k)
	asert.Equal(black, n.right.c)
	asert.Equal(10, n.right.left.k)
	asert.Equal(red, n.right.left.c)
	/*       6r
	       /   \
	    2b      xb          2r - 6b
	   / \     /           / \    \
	  1b 5b   11b         1b 5b  11b
	         /                   /
	       10r                  10r          */

	n.remove(6)
	asert.Equal(10, n.k)
	/*
	   2r - 10b      2r - 10b
	  / \    \      / \    \
	 1b 5b  11b    1b 5b   11b
	         /
	       xr                                 */

	n.remove(10)
	asert.Equal(5, n.k)
	asert.Equal(black, n.c)
	asert.Equal(11, n.right.k)
	asert.Equal(black, n.right.c)
	asert.Equal(2, n.left.k)
	asert.Equal(black, n.left.c)
	asert.Equal(1, n.left.left.k)
	asert.Equal(red, n.left.left.c)
	/*
	     2r - xb         5b
	    / \    \        / \
	   1b 5b  11b   1r-2b 11b       */

	n, _ = n.remove(11) //same cate with remove 5
	asert.Equal(2, n.k)
	asert.Equal(red, n.c)
	asert.Equal(5, n.right.k)
	asert.Equal(black, n.right.c)
	asert.Equal(1, n.left.k)
	asert.Equal(black, n.left.c)
	asert.Nil(n.left.left)
	asert.Nil(n.left.right)
	asert.Nil(n.right.left)
	asert.Nil(n.right.right)
	/*
	         5b        5b    2b     2r
	        / \        /    / \    / \
	   1r-2b  xb   1r-2r   1r 5r  1b 5b      */

	n, _ = n.remove(1)
	asert.Equal(5, n.k)
	asert.Equal(black, n.c)
	asert.Equal(2, n.left.k)
	asert.Equal(red, n.left.c)
	/*
	     2r        5b
	    / \        /
	  xb  5b      2r         */

	n.remove(5)
	asert.Equal(2, n.k)
	asert.Equal(black, n.c)

	n, _ = n.remove(2)
	asert.Nil(n)
}

func Test_rotateLeft(t *testing.T) {
	asert := assert.New(t)
	p := &rbnode{k: 3}
	n := &rbnode{k: 1}
	connectLeft(p, n)
	r := &rbnode{k: 2, c: red}
	a, b, c := &rbnode{v: "a"}, &rbnode{v: "b"}, &rbnode{v: "c"}
	n.right = r
	n.left, r.left, r.right = a, b, c

	n = rotateLeft(n)
	asert.Equal(2, n.k)
	asert.Equal(1, n.left.k)
	asert.Equal(a, n.left.left)
	asert.Equal(b, n.left.right)
	asert.Equal(c, n.right)
	asert.Equal(black, n.c)
	asert.Equal(red, n.left.c)
	asert.Equal(3, n.parent.k)
	asert.Equal(n, p.left)

	n = &rbnode{k: 1, c: red}
	r = &rbnode{k: 2}
	n.right = r
	n = rotateLeft(n)
	asert.Equal(red, n.c)
	asert.Equal(black, n.left.c)
}

func Test_rotateRight(t *testing.T) {
	asert := assert.New(t)
	p := &rbnode{k: 0}
	n := &rbnode{k: 2}
	connectRight(p, n)
	l := &rbnode{k: 1, c: red}
	a, b, c := &rbnode{v: "a"}, &rbnode{v: "b"}, &rbnode{v: "c"}
	n.left = l
	l.left, l.right, n.right = a, b, c

	n = rotateRight(n)
	asert.Equal(1, n.k)
	asert.Equal(2, n.right.k)
	asert.Equal(a, n.left)
	asert.Equal(b, n.right.left)
	asert.Equal(c, n.right.right)
	asert.Equal(black, n.c)
	asert.Equal(red, n.right.c)
	asert.Equal(0, n.parent.k)
	asert.Equal(n, p.right)

	n = &rbnode{k: 2, c: red}
	l = &rbnode{k: 1}
	n.left = l
	n = rotateRight(n)
	asert.Equal(red, n.c)
	asert.Equal(black, n.right.c)
}

func Test_connectRbnode(t *testing.T) {
	asert := assert.New(t)
	aTestRbTree := func() *rbnode {
		n := &rbnode{k: 5}
		l := &rbnode{k: 3}
		r := &rbnode{k: 8}
		n.left, n.right = l, r
		l.parent, r.parent = n, n
		return n
	}
	t.Run("left", func(t *testing.T) {
		n := aTestRbTree()
		connectLeft(n, nil)
		asert.Nil(n.left)

		l := &rbnode{v: "l"}
		connectLeft(n, l)
		asert.Equal("l", n.left.v)
		asert.Equal(5, l.parent.k)
	})

	t.Run("right", func(t *testing.T) {
		n := aTestRbTree()
		connectRight(n, nil)
		asert.Nil(n.right)

		r := &rbnode{v: "r"}
		connectRight(n, r)
		asert.Equal("r", n.right.v)
		asert.Equal(5, r.parent.k)
	})
}

func Test_swapInOrderSuccessorRb(t *testing.T) {
	asert := assert.New(t)
	t.Run("all black", func(t *testing.T) {
		n := &rbnode{k: 5}
		s := swapSuccessorRb(n)
		asert.Equal(n, s)

		l := &rbnode{k: 3}
		r := &rbnode{k: 8, v: "8"}
		n.left, n.right = l, r
		s = swapSuccessorRb(n)
		asert.Equal(8, n.k)
		asert.Equal(5, s.k)
		asert.Equal("8", n.v)

		rl := &rbnode{k: 7}
		r.left = rl
		s = swapSuccessorRb(n)
		asert.Equal(7, n.k)
	})

	t.Run("left red", func(t *testing.T) {
		n := &rbnode{k: 5}
		l := &rbnode{k: 3, c: red}
		r := &rbnode{k: 8}
		connectLeft(n, l)
		connectRight(n, r)
		s := swapSuccessorRb(n)
		asert.Equal(5, s.k)
		asert.Equal(3, n.k)

		ll := &rbnode{k: 1}
		lr := &rbnode{k: 4}
		connectLeft(l, ll)
		connectRight(l, lr)
		s = swapSuccessorRb(n)
		asert.Equal(8, n.k)
	})
}

type testRbTree struct {
	root       *rbnode
	a, b, c, d *rbnode
	pu         *rbnode
}

func Test_borrowDownwardRb(t *testing.T) {
	asert := assert.New(t)
	asertDistance := func(r *rbnode, ns ...*rbnode) {
		ln := len(ns)
		ds := make([]int, ln)
		for i, n := range ns {
			d := distanceRb(r, n, 0)
			ds[i] = d
		}
		d0 := ds[0]
		for i := 1; i < ln; i++ {
			asert.Equal(d0, ds[i])
		}
	}

	/*
	      5b - p
	    /   \
	   3b   8b
	   /\   /\
	  a b  c d              */
	aRbTree := func() testRbTree {
		n := &rbnode{k: 5, c: red}
		l := &rbnode{k: 3}
		r := &rbnode{k: 8}
		connectLeft(n, l)
		connectRight(n, r)
		a, b, c, d := &rbnode{v: "a"}, &rbnode{v: "b"}, &rbnode{v: "c"}, &rbnode{v: "d"}
		pu := &rbnode{v: "pu"}
		connectLeft(l, a)
		connectRight(l, b)
		connectLeft(r, c)
		connectRight(r, d)
		connectLeft(pu, n)
		return testRbTree{n, a, b, c, d, pu}
	}

	/*
	    [5b]          [xb]
	    /  \           /
	  (xb) 8b   (5r)-8b       (5r)-[8b]
	   /  / \   / \   \       / \    \
	  a  c  d  a  c   d      a  c    d   */
	t.Run("black L/ black parent", func(*testing.T) {
		t := aRbTree()
		p := t.root
		p.c = black
		n := p.left
		disconnectRb(p.parent, p)

		borrowDownwardRb(n, p)
		asert.Equal(8, p.k)
		asert.Equal(n, p.left)
		asert.Equal(5, n.k)
		asert.Equal(red, n.c)
		asert.Equal("a", n.left.v)
		asert.Equal("c", n.right.v)
		asert.Equal("d", p.right.v)
		asertDistance(p, t.a, t.c, t.d)
	})

	/*
		 [5b]          [xb]
		/   \           /
	   3b  (xb)   3r-(5b)       3r-[5b]
	   /\  /     / \   \       / \    \
	  a b c     a  b   c      a  b    c   */
	t.Run("black R/ black parent / black L", func(*testing.T) {
		t := aRbTree()
		p := t.root
		p.c = black
		n := p.right
		disconnectRb(p.parent, p)

		borrowDownwardRb(n, p)
		asert.Equal(5, p.k)
		asert.Equal(3, p.left.k)
		asert.Equal(red, p.left.c)
		asert.Equal(black, p.c)
		asert.Equal("a", p.left.left.v)
		asert.Equal("b", p.left.right.v)
		asert.Equal("c", p.right.v)
		asertDistance(p, t.a, t.b, t.c)
	})

	/*
		 [5b]       [bb]          [bb]
		/   \       /  \          /   \
	   3r  (xb)   3r  (5b)    ar-3b   (5b)
	   /\   |    / \   |          \   / \
	  ab b  c   ab  x  c          b1 b2  c    */
	t.Run("black R/ black parent / red L(black children)", func(*testing.T) {
		t := aRbTree()
		p := t.root
		p.c = black
		p.left.c = red
		n := p.right
		disconnectRb(t.pu, p)
		b1, b2 := &rbnode{v: "b1"}, &rbnode{v: "b2"}
		connectLeft(t.b, b1)
		connectRight(t.b, b2)

		borrowDownwardRb(n, p)
		asert.Equal(black, p.c)
		asert.Equal("b", p.v)
		asert.Equal(5, p.right.k)
		asert.Equal(black, p.right.c)
		asert.Equal(b2, p.right.left)
		asert.Equal(t.c, p.right.right)
		asert.Equal(3, p.left.k)
		asert.Equal(black, p.left.c)
		asert.Equal(t.a, p.left.left)
		asert.Equal(red, p.left.left.c)
		asert.Equal(b1, p.left.right)
	})

	/*
		 [5b]
		/   \
	   3r  (xb)
	   /\  /
	  ar b c        */
	t.Run("black R/ black parent / red L(red left child)", func(*testing.T) {
		// no such case
	})

	/*
	    [5r] - pu          pu
	    /  \              /
	  (xb)  8b    (5r)-[8b]
	   /   / \     / \    \
	  a   c  d     a c    d                   */
	t.Run("black L/ red parent", func(*testing.T) {
		t := aRbTree()
		p := t.root
		n := p.left
		borrowDownwardRb(n, p)
		asert.Equal(5, n.k)
		asert.Equal(red, n.c)
		asert.Equal("a", n.left.v)
		asert.Equal("c", n.right.v)
		asert.Equal("d", p.right.v)
		asert.Equal(black, p.c)
		asert.Equal(n, p.left)
		asert.Equal("pu", p.parent.v)
		asert.Equal(p, p.parent.left)
		asertDistance(p.parent, t.a, t.c, t.d)
	})

	/*
	     [5r] - pu          pu
	    /   \              /
	   3b   xb     3r -[5b]
	   /\  /      / \    \
	  a b c       a b    c                   */
	t.Run("black R/ red parent", func(*testing.T) {
		t := aRbTree()
		p := t.root
		n := p.right

		borrowDownwardRb(n, p)
		asert.Equal(5, p.k)
		asert.Equal(black, p.c)
		l := p.left
		asert.Equal(p, l.parent)
		asert.Equal(3, l.k)
		asert.Equal(red, l.c)
		asert.Equal("a", l.left.v)
		asert.Equal("b", l.right.v)
		asert.Equal("c", p.right.v)
		asert.Equal("pu", p.parent.v)
		asert.Nil(n.parent)
		asert.Nil(n.left)
		asert.Nil(n.right)
		asertDistance(p.parent, t.a, t.b, t.c)
	})

	/*
	    [5b]          [5b]
	    /  \         /   \
	  (xr)  8b      a    8b
	   /   / \           / \
	  a   c  d           c d    */
	t.Run("red L/ black parent", func(t *testing.T) {
		p := aRbTree().root
		n := p.left
		p.c, n.c = black, red

		borrowDownwardRb(n, p)
		asert.Equal("a", p.left.v)
	})

	/*
	    [5b]
	    /  \
	   3b  xr
	   /\  /
	  a b c                 */
	t.Run("red R/ black parent", func(t *testing.T) {
		// no such case
	})

	t.Run("red / red parent", func(t *testing.T) {
		// no such case
	})
}
