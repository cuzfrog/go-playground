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
	cap   int
	heap  []int
	items []interface{}
	//rndices []uint
}

/* ----- constructor ------*/

func newIndexMinPriorityQueue(cap int) (*indexMinPriorityQueue, error) {
	if cap <= 0 {
		return nil, fmt.Errorf("cap must be greater than 0")
	}
	return &indexMinPriorityQueue{
		cap,
		make([]int, 1, cap),
		make([]interface{}, cap),
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
		swimLast(pq.heap)
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
	panic("implement me")
}

func (pq *indexMinPriorityQueue) delete(k int) {
	panic("implement me")
}

func (pq *indexMinPriorityQueue) minIndex() int {
	panic("implement me")
}

func (pq *indexMinPriorityQueue) removeMin() interface{} {
	panic("implement me")
}
