package kendallTau

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"go-playground/sorting"
)

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
	} // auxA, auxB have been sorted

	for i := 0; i < n; i++ {
		if auxA[i] != auxB[i]{
			return fmt.Errorf("elem %v in a is not in b", auxA[i])
		}
	}
	return nil
}

func Test_countPair(t *testing.T)  {
	assert.Equal(t, 1, countPair(2))
	assert.Equal(t, 3, countPair(3))
	assert.Equal(t, 6, countPair(4))
}