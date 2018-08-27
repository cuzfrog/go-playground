package sorting

func mergesort(a []int) []int {
	n := len(a)
	if n > 1 {
		mid := n / 2
		a1 := a[:mid]
		a2 := a[mid+1:]
		return merge(mergesort(a1), mergesort(a2))
	} else {
		return a
	}
}

func merge(a1, a2 []int) []int {
	l1 := len(a1)
	l2 := len(a2)
	aux := make([]int, l1+l2)
	a := append(a1, a2...)
	copy(aux, a)

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
	return aux
}
