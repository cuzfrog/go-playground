package sorting

func quickSort(a []int) {
	if len(a) <= 1 {
		return
	}
	j := partition(a)
	l := a[:j]
	r := a[j+1:]
	quickSort(l)
	quickSort(r)
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
		exch(a, l, r)
	}
	exch(a, 0, r)
	return r
}

func quicksort3way(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}

	v := a[0]
	lt, i, gt := 0, 1, n-1
	for i <= gt {
		elm := a[i]
		if elm > v {
			exch(a, i, gt)
			gt--
		} else if elm < v {
			exch(a, lt, i)
			lt++
			i++
		} else {
			i++
		}
	}
	//fmt.Printf("Result lt=%v, gt=%v, a:%v\n", lt, gt, a)
	quicksort3way(a[:lt])
	quicksort3way(a[gt+1:])
}

/* ----------- exported ----------- */

// QuickSort generic implementation
func QuickSort(a []Comparable) {
	if len(a) <= 1 {
		return
	}
	j := partitionGeneric(a)
	l := a[:j]
	r := a[j+1:]
	QuickSort(l)
	QuickSort(r)
}

func partitionGeneric(a Comparables) int {
	n := len(a)
	if n <= 1 {
		return 0
	}
	e := a[0]
	l, r := 1, n-1
	for ; ; l, r = l+1, r-1 {
		for l < n && a[l].compare(&e) < 0 {
			l++
		}
		for a[r].compare(&e) > 0 {
			r--
		}
		if l >= r {
			break
		}
		a.Exch(l, r)
	}
	a.Exch(0, r)
	return r
}
