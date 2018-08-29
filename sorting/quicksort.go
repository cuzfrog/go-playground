package sorting

func quicksort(a []int) {
	if len(a) <= 1 {
		return
	}
	j := partition(a)
	l := a[:j]
	r := a[j+1:]
	quicksort(l)
	quicksort(r)
}

func partition(a []int) int {
	n := len(a)
	if n <= 1 {
		return 0
	}
	e := a[0]
	l, r := 1, n-1
	for ; ; l, r = l+1, r-1 {
		for l < n && a[l] < e {
			l++
		}
		for a[r] > e {
			r--
		}
		if l >= r {
			break
		}
		exchange(a, l, r)
	}
	exchange(a, 0, r)
	return r
}
