package sorting

import (
	"testing"
	"math/rand"
	"github.com/stretchr/testify/assert"
)

func test(f func([]int) []int, t *testing.T) {
	a := genElems(32)
	t.Log("Before:", a)
	s := f(a)
	t.Log("Sorted:", s)
	checkSorted(s, t)
}

func TestSorting(t *testing.T) {
	t.Run("insertion sort", func(t *testing.T) { test(insertionSort, t) })
}

func checkSorted(a []int, t *testing.T) {
	n := len(a)
	for i := 1; i < n; i++ {
		e1 := a[i-1]
		e2 := a[i]
		b := e1 <= e2
		assert.Truef(t, b, "at index %d, elem %d is smaller than its previous elem %d", i, e1, e2)
		if !b {
			t.FailNow()
		}
	}
}

func genElems(n int) []int {
	a := make([]int, n, n)
	max := n * 4
	for i := range a {
		a[i] = rand.Intn(max)
	}
	return a
}

func benchmark(f func([]int) []int, n int, b *testing.B) {
	a := genElems(n)
	as := make([][]int, b.N, b.N)
	for i := range as {
		as[i] = make([]int, n, n)
		copy(as[i], a)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(as[i])
	}
}

func Benchmark_InsertionSort(b *testing.B) {
	f := insertionSort
	b.Run("100", func(b *testing.B) { benchmark(f, 100, b) })
	b.Run("1000", func(b *testing.B) { benchmark(f, 1000, b) })
	b.Run("10000", func(b *testing.B) { benchmark(f, 10000, b) })
}
