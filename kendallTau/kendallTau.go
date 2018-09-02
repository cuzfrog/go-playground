package kendallTau

import (
	"sort"
)

/* Input Contract:
	0. a, b are permutations, so they have the same length and same elements(0, 1, 2, ..., N-1)
	1. no duplicates in input slice
	2. a, b both have length > 1
*/

// quadratic complexity
func kendallTauDistance1(a, b []int) int {
	pa := getParis(a)
	pb := getParis(b)
	if len(pa) != len(pb) {
		panic("lengths of pairs are different")
	}

	sort.Slice(pa, func(i, j int) bool {
		return pa[i].compare(&pa[j]) < 0
	})
	sort.Slice(pb, func(i, j int) bool {
		return pb[i].compare(&pb[j]) < 0
	})

	//fmt.Println("Pairs1:", pa)
	//fmt.Println("Pairs2:", pb)

	dis := 0
	for i := 0; i < len(pa); i++ {
		d := pa[i].distance(&pb[i])
		dis = dis + d
	}
	return dis
}

// quadratic complexity
func kendallTauDistance2(a, b []int) int {
	n := len(a)
	ai, bi := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		ai[a[i]], bi[b[i]] = i, i
	}

	dis := 0
	for i, j := 0, 1; i < n; {
		ao, bo := ai[i] > ai[j], bi[i] > bi[j]
		if ao != bo { //order different
			dis++
		}
		j++
		if j == n {
			i++
			j = i
		}
	}
	return dis
}

func kendallTauDistance3(a, b []int) int {
	n := len(a)
	ai := make([]int, n)
	for i := 0; i < n; i++ {
		ai[a[i]] = i
	}

	biToAi := make([]int, n) //b's index to a'index
	for i := range biToAi {
		biToAi[i] = ai[b[i]]
	}

	//aiToBi := make([]int, n)
	//for i := range aiToBi{
	//	aiToBi[biToAi[i]] = i
	//}
	return countInversion(biToAi, make([]int, n))
}

func countInversion(a []int, buf []int) int {
	n := len(a)
	if n <= 1 {
		return 0
	}
	if n == 2 {
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
			return 1
		} else {
			return 0
		}
	}
	mid := n / 2
	a1 := a[:mid]
	a2 := a[mid : n]
	c1, c2 := countInversion(a1, buf), countInversion(a2, buf)
	return c1 + c2 + mergeCntInv(a1, a2, buf)
}

func mergeCntInv(a, b []int, buf []int) int {
	na := len(a)
	nb := len(b)
	n := na + nb
	aux := buf[:n]
	var va, vb, dis int
	for k, ia, ib := 0, 0, 0; k < n; k++ {
		if ia >= na {
			aux[k] = b[ib]
			ib++
		} else if ib >= nb {
			aux[k] = a[ia]
			ia++
		} else {
			va, vb = a[ia], b[ib]
			if va < vb {
				aux[k] = va
				ia++
			} else if va > vb {
				aux[k] = vb
				ib++
				dis = dis + (na - ia)
			} else {
				panic("duplicate elem")
			}
		}
	}
	ab := a[:n]
	//fmt.Printf("a:%v, b:%v\n",a,b)
	copy(ab, aux)
	return dis
}
