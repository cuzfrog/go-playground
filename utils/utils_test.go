package utils

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func Test_hashDistributionForRandInt(t *testing.T) {
	M := 10000
	N := 10
	R := 0.05
	table := make([]int, N)
	for i := 0; i < M; i++ {
		table[Hash(rand.Int(), N)]++
	}
	sort.Ints(table)
	assert.True(t, Abs(table[0]-table[N-1]) < int(R*float64(M)))
}
