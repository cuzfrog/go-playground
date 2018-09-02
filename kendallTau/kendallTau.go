package kendallTau

import (
	"sort"
)

/* Input Contract:
	0. a, b are permutations, so they have the same length and same elements
	1. no duplicates in input slice
	2. a, b both have length > 1
*/

// quadratic complexity
func kendallTauDistance1(a, b []int) int {
	pa := getParis(a)
	pb := getParis(b)
	if len(pa) != len(pb){
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
