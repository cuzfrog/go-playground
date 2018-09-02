package kendallTau

import "fmt"

type pair struct {
	p  int
	q  int
	lo int
	hi int
}

func (p1 *pair) compare(p2 *pair) int8 {
	if p1.lo > p2.lo {
		return 1
	} else if p1.lo < p2.lo {
		return -1
	} else {
		if p1.hi > p2.hi {
			return 1
		} else if p1.hi < p2.hi {
			return -1
		} else {
			return 0
		}
	}
}

func (p1 *pair) distance(p2 *pair) int {
	if p1.p == p2.p && p1.q == p2.q {
		return 0 //equal
	} else if p1.p == p2.q && p1.q == p2.p {
		return 1 //dis
	} else {
		panic(fmt.Sprintf("Two pairs are irrelevant: %v, %v", *p1, *p2))
	}
}

func (p pair) String() string {
	return fmt.Sprintf("{%v, %v}", p.p, p.q)
}

func getParis(a []int) []pair {
	n := len(a)
	pa := make([]pair, countPair(n))
	for i, j, k := 0, 1, 0; i < n-1; {
		vi := a[i]
		vj := a[j]
		var p pair
		if vi > vj {
			p = pair{vi, vj, vj, vi}
		} else {
			p = pair{vi, vj, vi, vj}
		}

		pa[k] = p
		k++
		j++
		if j == n {
			i++
			j = i + 1
		}
	}
	return pa
}

func countPair(n int) int {
	return (n*n - n) / 2
}
