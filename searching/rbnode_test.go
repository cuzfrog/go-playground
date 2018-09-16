package searching

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_insert(t *testing.T) {
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
	n := &rbnode{k: 5}
	s := swapInOrderSuccessorRb(n)
	asert.Equal(n, s)

	l := &rbnode{k: 3}
	r := &rbnode{k: 8, v: "8"}
	n.left, n.right = l, r
	s = swapInOrderSuccessorRb(n)
	asert.Equal(8, n.k)
	asert.Equal(5, s.k)
	asert.Equal("8", n.v)

	rl := &rbnode{k: 7}
	r.left = rl
	s = swapInOrderSuccessorRb(n)
	asert.Equal(7, n.k)
}
