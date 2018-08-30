package sorting

// PriorityQueue a basic implementation
type PriorityQueue interface {
	size() int
	isEmpty() bool
	insert(elem int)
	popMax() int
	drop()
}

/* ----- definition ------*/

type heapPriorityQueue struct {
	heap []int
}

/* ----- constructor ------*/

// New creates a new PriorityQueue from a slice of elements
func New(a []int) *heapPriorityQueue {
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
	swimLast(pq.heap)
}

func (pq *heapPriorityQueue) popMax() (max int) {
	h := pq.heap
	max = h[1]
	k := len(h) - 1
	h[1] = h[k]
	sink(h, 1)
	pq.heap = h[:k]
	return
}

func (pq *heapPriorityQueue) drop() {
	pq.heap = pq.heap[:len(pq.heap)-1]
}

/* ----- private methods ------*/

// swimLast heapifies a heap with a newly added elem.
// Contract: Without the elem, the heap's property holds.
func swimLast(h []int) {
	i := len(h) - 1
	p := i / 2
	for i > 1 && h[p] < h[i] {
		Exch(h, p, i)
		i = p
		p = i / 2
	}
}

// sink heapifies a heap by moving a node down to proper its proper place
// return the node value
func sink(h []int, k int) (v int) {
	n := len(h)
	if k >= n {
		panic("out of index")
	}
	v = h[k]
	for k*2 < n { //if one of the children exists
		s := k*2 + 1
		if s >= n || h[s] < h[s-1] {
			s--
		}
		if h[s] > h[k] {
			Exch(h, s, k)
			k = s
		} else {
			break
		}
	}
	return
}
