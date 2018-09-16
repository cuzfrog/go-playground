package sorting

import (
	"math/rand"
	"fmt"
)

func exch(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
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

// GenNonDupElems generates a new slice with n elements with no duplicate
func GenNonDupElems(n int) []int {
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = i*5 - rand.Intn(4)
	}
	for i := range a {
		j := rand.Intn(n)
		a[i], a[j] = a[j], a[i]
	}
	return a
}

type compare func(int, int) bool

// swimf heapifies a heap with a newly added elem.
// return final index
// Contract: Without the elem, the heap's property holds.
func swimf(h []int, i int, comp compare, hi []int) int {
	n := len(h)
	if i >= n {
		panic("out of index")
	}
	p := i / 2
	for i > 1 && comp(h[p], h[i]) {
		exch(h, p, i)
		if hi != nil {
			kp, ki := h[p], h[i]
			exch(hi, kp, ki)
		}
		i = p
		p = i / 2
	}
	return i
}

func swimMax(h []int, i int, hi []int) int {
	return swimf(h, i, func(a, b int) bool { return a < b }, hi)
}

func swimMin(h []int, i int, hi []int) int {
	return swimf(h, i, func(a, b int) bool { return a > b }, hi)
}

func swimMaxLast(h []int) int {
	return swimMax(h, len(h)-1, nil)
}

func swimMinLast(h []int) int {
	return swimMin(h, len(h)-1, nil)
}

// sinkf heapifies a heap by moving a node down to proper its proper place
// return the node value and index on the heap
func sinkf(h []int, i int, comp compare, hi []int) (int, int) {
	n := len(h)
	if i < 0 || i >= n {
		panic(fmt.Sprintf("i:%v is out of bound:[0,%v)", i, n))
	}
	v := h[i]
	for i*2 < n { //if one of the children exists
		s := i*2 + 1
		if s >= n || comp(h[s-1], h[s]) {
			s--
		}
		if comp(h[s], h[i]) {
			exch(h, s, i)
			if hi != nil {
				ks, ki := h[s], h[i]
				exch(hi, ks, ki)
			}
			i = s
		} else {
			break
		}
	}
	return v, i
}

func sinkMax(h []int, i int, hi []int) (int, int) {
	return sinkf(h, i, func(a, b int) bool { return a > b }, hi)
}

func sinkMin(h []int, i int, hi []int) (int, int) {
	return sinkf(h, i, func(a, b int) bool { return a < b }, hi)
}

func sinkMaxFirst(h []int) (int, int) {
	return sinkMax(h, 1, nil)
}

func sinkMinFirst(h []int) (int, int) {
	return sinkMin(h, 1, nil)
}
