package sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_indexMinPQ(t *testing.T) {
	pq, _ := newIndexMinPriorityQueue(16)
	assert.Equal(t, 0 , pq.size())
	assert.True(t, pq.isEmpty())

	k := 3
	pq.put(k, "5")
	assert.Equal(t, []int{0, 3}, pq.heap)
	assert.Equal(t, "5", pq.items[k])
	assert.Equal(t, 1, pq.indices[3])
	k = 2
	pq.put(k, 5)
	assert.Equal(t, []int{0, 2, 3}, pq.heap)
	k = 5
	pq.put(k, 13)
	assert.Equal(t, []int{0, 2, 3, 5}, pq.heap)
	assert.Equal(t, 1, pq.indices[2])
	assert.Equal(t, 2, pq.indices[3])
	assert.Equal(t, -1, pq.indices[4])
	assert.Equal(t, 3, pq.indices[5])

	assert.Equal(t, 2, pq.minIndex())
	assert.Equal(t, 5, pq.min())

	assert.Equal(t, 5, pq.removeMin())
	assert.Equal(t, nil, pq.items[2])
	assert.Equal(t, []int{0, 3, 5}, pq.heap)
	assert.Equal(t, -1, pq.indices[2])
	assert.Equal(t, 1, pq.indices[3])
	assert.Equal(t, 2, pq.indices[5])

	assert.False(t, pq.isEmpty())
	assert.Equal(t, 2, pq.size())
}
