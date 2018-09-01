package sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_indexMinPQ(t *testing.T) {
	pq, _ := newIndexMinPriorityQueue(16)

	k := 3
	pq.put(k, "5")
	assert.Equal(t, []int{0, 3}, pq.heap)
	assert.Equal(t, "5", pq.items[k])
	k = 2
	pq.put(k, 5)
	assert.Equal(t, []int{0, 2, 3}, pq.heap)
	k = 5
	pq.put(k, 13)
	assert.Equal(t, []int{0, 2, 3, 5}, pq.heap)

	assert.Equal(t, 2, pq.minIndex())
	assert.Equal(t, 5, pq.min())

	assert.Equal(t, 5, pq.removeMin())
	assert.Equal(t, nil, pq.items[2])
	assert.Equal(t, []int{0, 3, 5}, pq.heap)
}
