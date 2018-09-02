package sorting

func dedup(a []int) []int {
	quickSort(a)
	n := len(a)
	d := make([]int, n, n)
	d[0] = a[0]
	i := 1
	for j := 1; j < n; j++ {
		if a[j] != d[i-1] {
			d[i] = a[j]
			i++
		}
	}
	return d[:i]
}
