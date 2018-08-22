package sorting

func insertionSort(a []int) []int {
	n := len(a)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			exchange(a, j, j - 1)
		}
	}
	return a
}
