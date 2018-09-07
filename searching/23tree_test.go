package searching

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_splitRootNode3(t *testing.T) {
	n := &node3{el: entry{2, "s"}, er: entry{5, "gg"}}
	l := &entry{1, "left entry"}
	m := &entry{3, "mid entry"}
	r := &entry{8, "right entry"}

	ltree := splitRootNode3(n, l)
	assert.Equal(t, 2, ltree.e.k)

	lltree, ok := ltree.left.(*node2)
	assert.True(t, true, ok)
	assert.Equal(t, 1, lltree.e.k)

	lrtree := ltree.right.(*node2)
	assert.Equal(t, 5, lrtree.e.k)

	mtree := splitRootNode3(n, m)
	assert.Equal(t, 3, mtree.e.k)

	rtree := splitRootNode3(n, r)
	assert.Equal(t, 5, rtree.e.k)
}
