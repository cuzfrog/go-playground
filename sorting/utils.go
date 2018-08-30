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