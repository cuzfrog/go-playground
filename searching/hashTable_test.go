package searching

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_put(t *testing.T) {
	asert := assert.New(t)

	ht := hashTable{make([]*linkedNode, 10), 0}

	ht.put(1, "1")
	asert.Equal(1, ht.cnt)
	asert.Equal("1", ht.get(1))

	ht.put(1, "11")
	asert.Equal(1, ht.cnt)
	asert.Equal("11", ht.get(1))

	ht.put(3, "3")
	asert.Equal(2, ht.cnt)
	asert.Equal("3", ht.get(3))

	for i := 0; i < 10000; i++ {
		ht.put(i, string(i))
	}
	asert.Equal(10000, ht.cnt)
}
