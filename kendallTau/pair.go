package kendallTau

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
		return -1 //irrelevant
	}
}
