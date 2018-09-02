package utils

import "math/rand"

func IntMin(l, r int) (m int) {
	if l < r {
		m = l
	} else {
		m = r
	}
	return
}

func IntMax(l, r int) (m int) {
	if l > r {
		m = l
	} else {
		m = r
	}
	return
}

func Suffle(a []int) {
	n := len(a)
	var j int
	for i := range a {
		j = rand.Intn(n)
		a[i], a[j] = a[j], a[i]
	}
}
