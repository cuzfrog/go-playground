package kendallTau

import (
	"sort"
	"fmt"
)

/* Input Contract:
	0. a, b are permutations, so they have the same length and same elements
	1. no duplicates in input slice
	2. a, b both have length > 1
*/

// quadratic complexity
func kendallTauDistance1(a, b []int) int {
	n := len(a)
	pa := getParis(a)
	pb := getParis(b)

	sort.Slice(pa, func(i, j int) bool {
		return pa[i].compare(&pa[j]) < 0
	})
	sort.Slice(pb, func(i, j int) bool {
		return pb[i].compare(&pb[j]) < 0
	})

	dis := 0
	for i := 0; i < n; i++ {
		d := pa[i].distance(&pb[i])
		if d<0{
			panic(fmt.Sprintf("Two pairs are irrelevant: %v, %v", pa[i], pb[i]))
		}
		dis = dis + d
	}
	return dis
}

func countPair(n int) int {
	return (n*n - n) / 2
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
		j++
		if j == n {
			i++
			j = i + 1
		}
	}
	return pa
}
