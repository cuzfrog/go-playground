package sorting

import "math/rand"

// Exch swaps two elements of slice 'a' at indices 'i' and 'j'
func Exch(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func min(l, r int) (m int) {
	if l < r {
		m = l
	} else {
		m = r
	}
	return
}

func max(l, r int) (m int) {
	if l > r {
		m = l
	} else {
		m = r
	}
	return
}

// GenElems generates a new slice with n elements
func GenElems(n int) []int {
	a := make([]int, n, n)
	max := n * 4
	for i := range a {
		a[i] = rand.Intn(max)
	}
	return a
}

type compare func(int, int) bool

// swimf heapifies a heap with a newly added elem.
// return final index
// Contract: Without the elem, the heap's property holds.
func swimf(h []int, i int, f compare) int {
	n := len(h)
	if i >= n {
		panic("out of index")
	}
	p := i / 2
	for i > 1 && f(h[p], h[i]) {
		Exch(h, p, i)
		i = p
		p = i / 2
	}
	return i
}

func swimMax(h []int, i int) int {
	return swimf(h, i, func(a, b int) bool { return a < b })
}

func swimMin(h []int, i int) int {
	return swimf(h, i, func(a, b int) bool { return a > b })
}

func swimMaxLast(h []int) int {
	return swimMax(h, len(h)-1)
}

func swimMinLast(h []int) int {
	return swimMin(h, len(h)-1)
}

// sinkf heapifies a heap by moving a node down to proper its proper place
// return the node value and index on the heap
func sinkf(h []int, i int, f compare) (int, int) {
	n := len(h)
	if i >= n {
		panic("out of index")
	}
	v := h[i]
	for i*2 < n { //if one of the children exists
		s := i*2 + 1
		if s >= n || f(h[s-1], h[s]) {
			s--
		}
		if f(h[s], h[i]) {
			Exch(h, s, i)
			i = s
		} else {
			break
		}
	}
	return v, i
}

func sinkMax(h []int, i int) (int, int) {
	return sinkf(h, i, func(a, b int) bool { return a > b })
}

func sinkMin(h []int, i int) (int, int) {
	return sinkf(h, i, func(a, b int) bool { return a < b })
}

func sinkMaxFirst(h []int) (int, int) {
	return sinkMax(h, 1)
}

func sinkMinFirst(h []int) (int, int) {
	return sinkMin(h, 1)
}