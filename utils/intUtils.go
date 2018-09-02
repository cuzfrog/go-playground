package utils

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