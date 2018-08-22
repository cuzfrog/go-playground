package sorting

func selectionSort(a []int) []int {
	n := len(a)
	minIdx := 0
	for i := range a {
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		exchange(a, i, minIdx)
	}
	return a
}

func insertionSort(a []int) []int {
	n := len(a)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			exchange(a, j, j-1)
		}
	}
	return a
}

func shellSort(a []int) []int {
	n := len(a)
	h := 1
	for h < n/3 {
		h = h*3 + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && a[j] < a[j-h]; j-=h {
				exchange(a, j, j-h)
			}
		}
		h = h / 3
	}
	return a
}
