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
	cap  int
	heap []int
}

/* ----- constructor ------*/

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
	reheapifyLast(pq.heap)
}

func (pq *heapPriorityQueue) popMax() (max int) {
	h := pq.heap
	max = h[1]

	//todo
	return
}

// reheapifyLast heapifies a heap with a newly added elem.
// Contract: Without the elem, the heap's property holds.
func reheapifyLast(h []int) {
	i := len(h) - 1
	p := i / 2
	for p > 0 && h[p] < h[i] {
		exchange(h, p, i)
		i = p
		p = i / 2
	}
}
