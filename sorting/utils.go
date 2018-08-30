package sorting

func exchange(a []int, i, j int) {
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