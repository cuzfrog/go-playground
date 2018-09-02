package kendallTau

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"go-playground/sorting"
	"sort"
)

func Test_kendallTauDistance(t *testing.T) {
	t.Run("impl1(quadratic)", func(t *testing.T) { test(kendallTauDistance1, t) })
}

func test(f func([]int, []int) int, t *testing.T) {
	a := []int{0, 3, 1, 6, 2, 5, 4}
	b := []int{1, 0, 3, 6, 4, 2, 5}
	d := f(a, b)
	assert.Equal(t, 4, d)

	a = []int{1, 2, 3, 4, 5}
	b = []int{3, 4, 1, 5, 2}
	d = f(a, b)
	assert.Equal(t, 5, d)

	sort.Ints(b)
	d = f(a, b)
	assert.Equal(t, 0, d)
}

func contractGuard(a, b []int) error {
	n := len(a)
	if n != len(b) {
		return fmt.Errorf("length of a:%v != length of b:%v", n, len(b))
	}
	if n <= 1 {
		return fmt.Errorf("length must be at least 2")
	}
	auxA := make([]int, n)
	auxB := make([]int, n)
	copy(auxA, a)
	copy(auxB, b)

	if len(sorting.Dedup(auxA)) != n || len(sorting.Dedup(auxB)) != n {
		return fmt.Errorf("there're duplicates in input slice")
	} // auxA, auxB have been sorted in Dedup

	for i := 0; i < n; i++ {
		if auxA[i] != auxB[i] {
			return fmt.Errorf("elem %v in a is not in b", auxA[i])
		}
	}
	return nil
}

func Test_countPair(t *testing.T) {
	assert.Equal(t, 1, countPair(2))
	assert.Equal(t, 3, countPair(3))
	assert.Equal(t, 6, countPair(4))
}

func Test_getParis(t *testing.T) {
	a := []int{3, 6, 4}
	p := getParis(a)
	assert.Equal(t, 3, len(p))

	assert.Equal(t, pair{3, 6, 3, 6}, p[0])
	assert.Equal(t, pair{3, 4, 3, 4}, p[1])
	assert.Equal(t, pair{6, 4, 4, 6}, p[2])
}

/* --------- benchmark --------- */

