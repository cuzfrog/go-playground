package sorting

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