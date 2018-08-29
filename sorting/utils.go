package sorting

func exchange(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func min(l, r int) int {
	if l < r {
		return l
	} else {
		return r
	}
}
