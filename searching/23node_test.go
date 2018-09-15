package searching

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_23node_functional(t *testing.T) {
	asert := assert.New(t)

	r := newNode2(23, "23")
	asert.Nil(r.getVal(1))
	asert.Equal("23", r.getVal(23))
	/*
	   23
	 */

	r.putVal(12, "12")
	asert.True(r.is3)
	asert.Equal(12, r.e.k)
	asert.Equal(23, r.er.k)
	asert.Equal("12", r.getVal(12))
	asert.Nil(r.getVal(100))
	/*
	   12|23
	 */

	r.putVal(15, "15")
	asert.False(r.is3)
	asert.Equal(15, r.e.k)
	asert.Equal(12, r.left.e.k)
	asert.Equal(23, r.right.e.k)
	/*
	     15
	     /\
	   12 23
	 */

	r.putVal(10, "10")
	asert.True(r.left.is3)
	asert.Equal(10, r.left.e.k)
	asert.Equal(12, r.left.er.k)
	/*
	     15
	    / \
	10|12 23
	 */

	r.putVal(11, "11")
	asert.True(r.is3)
	asert.False(r.left.is3)
	asert.Equal(11, r.e.k)
	asert.Equal(15, r.er.k)
	asert.Equal(10, r.left.e.k)
	asert.Equal(12, r.mid.e.k)
	/*
	     11|15
	    /  |  \
	  10  12  23
	 */

	r.putVal(13, "13")
	r.putVal(14, "14")
	asert.False(r.is3)
	asert.Equal(11, r.left.e.k)
	asert.Equal(15, r.right.e.k)
	asert.Equal(10, r.left.left.e.k)
	asert.Equal(12, r.left.right.e.k)
	asert.Equal(14, r.right.left.e.k)
	asert.Equal(23, r.right.right.e.k)
	/*     13
	      /  \
	     11  15
	    / \  / \
	  10 12 14  23
	 */

	r.putVal(14, "14a")
	asert.Equal("14a", r.getVal(14))

	v := r.removeVal(100)
	asert.Nil(v)

	v = r.removeVal(11)
	asert.Equal("11", v)
	asert.True(r.is3)
	asert.True(r.left.is3)
	asert.Equal(13, r.e.k)
	asert.Equal(15, r.er.k)
	asert.Equal(10, r.left.e.k)
	asert.Equal(12, r.left.er.k)
	/*    13
		 /  \
		12  15          13|15
	   / \  / \       /   |   \
	 10  x 14  23  10|12 14   23
	*/

	r.removeVal(13)
	asert.False(r.left.is3)
	/*
		 12|15
	   /   |   \
	 10   14   23
	*/

	r.removeVal(12)
	asert.True(r.left.is3)
	asert.Equal(15, r.e.k)
	asert.Equal(10, r.left.e.k)
	asert.Equal(14, r.left.er.k)
	asert.Equal(23, r.right.e.k)
	/*
		 15
	   /   \
	10|14  23
	*/

	r.removeVal(15)
	asert.Equal(14, r.e.k)
	asert.False(r.left.is3)
	/*
		 14
	   /   \
	 10    23
	*/

	r.removeVal(10)
	asert.True(r.is3)
	asert.Equal(23, r.er.k)
	/* 14|23  */
}
