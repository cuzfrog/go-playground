package sorting

import "fmt"

// IndexMinPriorityQueue a basic impl
type IndexMinPriorityQueue interface {
	size() int
	isEmpty() bool
	min() interface{}
	put(k int, v interface{}) error
	contains(k int) bool
	delete(k int)
	minIndex() int
	removeMin() interface{}
}

/* ----- definition ------*/

type indexMinPriorityQueue struct {
	cap     int
	heap    []int
	items   []interface{}
	indices []int
}

/* ----- constructor ------*/

func newIndexMinPriorityQueue(cap int) (*indexMinPriorityQueue, error) {
	if cap <= 0 {
		return nil, fmt.Errorf("cap must be greater than 0")
	}
	indices := make([]int, cap)
	for i := range indices {
		indices[i] = -1
	}

	return &indexMinPriorityQueue{
		cap,
		make([]int, 1, cap),
		make([]interface{}, cap),
		indices,
	}, nil
}

/* ----- method implementation ------*/

func (pq *indexMinPriorityQueue) isEmpty() bool {
	return pq.size() == 0
}

func (pq *indexMinPriorityQueue) put(k int, v interface{}) error {
	if k >= pq.cap {
		return fmt.Errorf("k:%v is out of index bound of cap:%v", k, pq.cap)
	}
	if pq.items[k] == nil { //insert
		pq.heap = append(pq.heap, k)
		t := len(pq.heap) - 1
		pq.indices[k] = t
		swimMin(pq.heap, t, pq.indices)
	}
	pq.items[k] = v
	return nil
}

func (pq *indexMinPriorityQueue) contains(k int) bool {
	return k < pq.cap && pq.items[k] != nil
}

func (pq *indexMinPriorityQueue) size() int {
	s := len(pq.heap) - 1
	if s < 0 {
		s = 0
	}
	return s
}

func (pq *indexMinPriorityQueue) min() interface{} {
	return pq.items[pq.minIndex()]
}

func (pq *indexMinPriorityQueue) delete(k int) {
	panic("implement me")
}

func (pq *indexMinPriorityQueue) minIndex() int {
	return pq.heap[1]
}

func (pq *indexMinPriorityQueue) removeMin() interface{} {
	minIdx := pq.minIndex()
	minItem := pq.items[minIdx]
	pq.items[minIdx] = nil

	h := pq.heap
	t := len(h) - 1
	k1, kt := h[1], h[t]
	h[1] = h[t]

	hi := pq.indices
	hi[k1], hi[kt] = -1, 1 //remove index of minKey, change kt's index to 1
	h = h[:t]
	sinkMin(h, 1, hi)

	pq.heap = h
	pq.indices = hi
	return minItem
}
