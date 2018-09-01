package sorting

// allocate aux array on every call
func mergeSort(a []int) {
	n := len(a)
	if n > 1 {
		mid := n / 2
		a1 := a[:mid]
		a2 := a[mid:]
		mergeSort(a1)
		mergeSort(a2)
		merge(a1, a2)
	}
}

// use a buf to avoid make new arr
func mergesortBuf(a []int) {
	buf := make([]int, len(a))
	mergeSortWithBuffer(a, buf)
}

func mergeSortWithBuffer(a []int, buf []int) {
	n := len(a)
	if n > 1 {
		mid := n / 2
		a1 := a[:mid]
		a2 := a[mid:]
		mergeSortWithBuffer(a1, buf)
		mergeSortWithBuffer(a2, buf)
		mergeWithBuffer(a1, a2, buf)
	}
}

func merge(a1, a2 []int) {
	l1 := len(a1)
	l2 := len(a2)
	aux := make([]int, l1+l2)
	mergeWithBuffer(a1, a2, aux)
}

func mergeWithBuffer(a1, a2 []int, buf []int) {
	l1 := len(a1)
	l2 := len(a2)
	n := l1 + l2
	aux := buf[:n]

	i1 := 0
	i2 := 0
	for j := range aux {
		if i1 >= l1 {
			aux[j] = a2[i2]
			i2++
		} else if i2 >= l2 {
			aux[j] = a1[i1]
			i1++
		} else if a1[i1] < a2[i2] {
			aux[j] = a1[i1]
			i1++
		} else {
			aux[j] = a2[i2]
			i2++
		}
	}

	a := a1[:n] //contract: a1, a2 are consecutive
	copy(a, aux)
}

func mergeSortFromBottom(a []int) {
	n := len(a)
	b := make([]int, n)
	copy(b, a)

	for sz := 1; sz < n; sz *= 2 {
		for i := 0; i < n-sz; i += sz * 2 {
			mid := i + sz
			end := min(i+sz*2, n)
			a1 := a[i:mid]
			a2 := a[mid:end]
			mergeWithBuffer(a1, a2, b)
		}
	}
}
