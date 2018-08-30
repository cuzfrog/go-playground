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
	swimLast(pq.heap)
}

func (pq *heapPriorityQueue) popMax() (max int) {
	h := pq.heap
	max = h[1]

	//todo
	return
}

// swimLast heapifies a heap with a newly added elem.
// Contract: Without the elem, the heap's property holds.
func swimLast(h []int) {
	i := len(h) - 1
	p := i / 2
	for i > 1 && h[p] < h[i] {
		exchange(h, p, i)
		i = p
		p = i / 2
	}
}

func sink(h []int, k int) {
	n := len(h)
	if k >= n {
		panic("out of index")
	}
	for k*2 < n { //if one of the children exists
		s := k*2 + 1
		if s >= n || h[s] < h[s-1] {
			s--
		}
		if h[s] > h[k] {
			exchange(h, s, k)
			k = s
		} else {
			break
		}
	}
}
