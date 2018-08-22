package sorting

import (
	"testing"
	"math/rand"
	"github.com/stretchr/testify/assert"
)

func checkSorted(a []int, t *testing.T) {
	n := len(a)
	for i := 1; i < n; i++ {
		e1 := a[i-1]
		e2 := a[i]
		b := e1 < e2
		assert.Truef(t, b, "at index %d, elem %d is smaller than its previous elem %d", i, e1, e2)
		if !b {
			t.FailNow()
		}
	}
}

func genElems(n int) []int {
	a := make([]int, n, n)
	for i := range a {
		a[i] = rand.Intn(2048)
	}
	return a
}

func TestInsertionSort(t *testing.T) {
	a := genElems(1000)
	insertionSort(a)
	checkSorted(a, t)
}
