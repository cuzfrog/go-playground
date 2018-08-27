package sorting

import "math"

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
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				exchange(a, j, j-h)
			}
		}
		h = h / 3
	}
	return a
}

func shellSort2(a []int) []int {
	n := len(a)
	l := int(math.Log(float64(n))/math.Log(3) + 1)
	hs := make([]int, l)
	hs[l-1] = 1
	for i := l-1; i > 0; i-- {
		hs[i-1] = hs[i]*3 + 1
	}
	for _, h := range hs {
		for i := h; i < n; i++ {
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				exchange(a, j, j-h)
			}
		}
	}
	return a
}
