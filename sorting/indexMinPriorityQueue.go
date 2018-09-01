package sorting

import "fmt"

// IndexMinPriorityQueue a basic impl
type IndexMinPriorityQueue interface {
	size() int
	isEmpty() bool
	min() interface{}
	put(k int, v interface{}) (interface{}, error)
	contains(k int) bool
	delete(k int) interface{}
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

// put inserts or updates an item with associated key 'k'.
//  return: (previous value, error)
func (pq *indexMinPriorityQueue) put(k int, v interface{}) (interface{}, error) {
	if k <= 0 || k >= pq.cap {
		return nil, fmt.Errorf("k:%v is out of index bound:[0,%v)", k, pq.cap)
	}
	if pq.items[k] == nil { //insert
		pq.heap = append(pq.heap, k)
		t := len(pq.heap) - 1
		pq.indices[k] = t
		swimMin(pq.heap, t, pq.indices)
	}
	prev := pq.items[k]
	pq.items[k] = v
	return prev, nil
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

func (pq *indexMinPriorityQueue) delete(k int) (interface{}, error) {
	if k < 0 || k >= pq.cap {
		return nil, fmt.Errorf("k:%v is out of bound:[0,%v)", k, pq.cap)
	}
	if pq.indices[k] < 0 {
		return nil, nil
	}

	i := pq.indices[k]
	h := pq.heap
	v := pq.items[k]

	t := len(h) - 1
	h[i] = h[t]
	pq.indices[h[t]] = i
	sinkMin(h, i, pq.indices)
	h = h[:t]

	pq.items[k] = nil
	pq.indices[k] = -1
	pq.heap = h

	return v, nil
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
	sinkMin(h, 1, hi)
	h = h[:t]

	pq.heap = h
	pq.indices = hi
	return minItem
}
