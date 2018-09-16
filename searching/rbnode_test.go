package searching

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_rotateLeft(t *testing.T) {
	asert := assert.New(t)
	n := &rbnode{k: 1}
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
}
