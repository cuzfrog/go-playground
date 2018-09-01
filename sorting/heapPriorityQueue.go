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