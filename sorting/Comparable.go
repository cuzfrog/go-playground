package sorting

// Comparable stands for any value that could be compared.
type Comparable interface {
	compare(that *Comparable) byte
}

// Comparables stands for a slice of Comparable
type Comparables []Comparable

// Exch swaps two elements of slice 'a' at indices 'i' and 'j'
func (ap *Comparables) Exch(i, j int) {
	a := *ap
	a[i], a[j] = a[j], a[i]
}

/* ---------- internal implementation ---------- */

// Int is alias of int
type Int int

func (v1 *Int) compare(v2 *Int) int8 {
	if *v1 > *v2 {
		return 1
	} else if *v1 < *v2 {
		return -1
	} else {
		return 0
	}
}
