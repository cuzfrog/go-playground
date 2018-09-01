package sorting

func heapSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}

	for k := n / 2; k >= 1; k-- {
		sink(a, k)
	}
	for i := n - 1; i >= 0; i++ {
		Exch(a, 0, i)

	}
}

// sink tries to heapify an array 'a' by sinking elem at index 'k-1'.
//  Contract: 1 <= k <= len(a)/2, so node a[k-1] has at least one child.
func sink(a []int, k int) {
	n := len(a)
	s := k*2 + 1
	for s-1 <= n { //at least left node exists
		if s > n || a[s-2] > a[s-1] { // only left child || left > right
			s--
		}
		if a[s-1] > a[k-1] {
			Exch(a, k-1, s-1)
			k = s
			s = k*2 + 1
		} else {
			break
		}
	}
}
