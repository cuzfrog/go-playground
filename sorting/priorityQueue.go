package sorting

// PriorityQueue a basic implementation
type PriorityQueue interface {
	size() int
	isEmpty() bool
	insert(elem int)
	popMax() int
}

/* ----- definition ------*/

type heapPriorityQueue struct {
	heap []int
}

/* ----- constructor ------*/

// NewHeapPriorityQueue creates a new PriorityQueue from a slice of elements
func NewHeapPriorityQueue(a []int) *heapPriorityQueue {
	pq := heapPriorityQueue{make([]int, 1, len(a)+1)}
	for _, elem := range a {
		pq.insert(elem)
	}
	return &pq
}

/* ----- method implementation ------*/

func (pq *heapPriorityQueue) size() (s int) {
	s = len(pq.heap) - 1
	if s < 0 {
		s = 0
	}
	return
}

func (pq *heapPriorityQueue) isEmpty() bool {
	return pq.size() == 0
}

func (pq *heapPriorityQueue) insert(elem int) {
	pq.heap = append(pq.heap, elem)
	swim(pq.heap, uint(len(pq.heap)-1))
}

func (pq *heapPriorityQueue) popMax() (max int) {
	h := pq.heap
	max = h[1]
	k := len(h) - 1
	h[1] = h[k]
	sinkFirst(h)
	pq.heap = h[:k]
	return
}

/* ----- private methods ------*/

// swim heapifies a heap with a newly added elem.
// return final index
// Contract: Without the elem, the heap's property holds.
func swim(h []int, i uint) uint {
	p := i / 2
	for i > 1 && h[p] < h[i] {
		Exchu(h, p, i)
		i = p
		p = i / 2
	}
	return i
}

func swimLast(h []int) uint {
	return swim(h, uint(len(h)-1))
}

// sink heapifies a heap by moving a node down to proper its proper place
// return the node value and index on the heap
func sink(h []int, i uint) (int, uint) {
	n := uint(len(h))
	if i >= n {
		panic("out of index")
	}
	v := h[i]
	for i*2 < n { //if one of the children exists
		s := i*2 + 1
		if s >= n || h[s] < h[s-1] {
			s--
		}
		if h[s] > h[i] {
			Exchu(h, s, i)
			i = s
		} else {
			break
		}
	}
	return v, i
}

func sinkFirst(h []int) (int, uint) {
	return sink(h, 1)
}

/* ----------------------------------------------------------------------------- */

// IndexMinPriorityQueue a basic impl
type IndexMinPriorityQueue interface {
	size() int
	isEmpty() bool
	min() int
	insertOrUpdate(k uint, v int)
	contains(k uint) bool
	delete(k uint)
	minIndex() uint
	removeMin() int
}

type indexMinPriorityQueue struct {
	cap     uint
	heap    []int
	indices []uint
	rndices []uint
}

func (pq *indexMinPriorityQueue) isEmpty() bool {
	return pq.size() == 0
}

func (pq *indexMinPriorityQueue) min() int {
	panic("implement me")
}

func (pq *indexMinPriorityQueue) insertOrUpdate(k uint, v int) {
	h := pq.heap
	hi := pq.indices[k]
	var fhi uint
	if hi > 0 {
		prev := h[hi]
		h[hi] = v
		if v > prev {
			fhi = swim(h, hi)
		} else {
			_, fhi = sink(h, hi)
		}
	} else {
		h = append(h, v)
		fhi = swimLast(h)
	}
	pq.indices[k] = fhi
}

func (pq *indexMinPriorityQueue) contains(k uint) bool {
	return pq.indices[k] > 0
}

func (pq *indexMinPriorityQueue) delete(k uint) {
	panic("implement me")
}

func (pq *indexMinPriorityQueue) minIndex() uint {
	panic("implement me")
}

func (pq *indexMinPriorityQueue) removeMin() int {
	panic("implement me")
}

func (pq *indexMinPriorityQueue) size() int {
	s := len(pq.heap) - 1
	if s < 0 {
		s = 0
	}
	return s
}
