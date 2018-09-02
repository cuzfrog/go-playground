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
		dis = dis + d
	}
	return dis
}
