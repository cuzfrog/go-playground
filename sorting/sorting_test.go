package sorting

import (
	"testing"
	"math/rand"
	"github.com/stretchr/testify/assert"
)

func test(f func([]int), t *testing.T) {
	for n := 0; n < 100; n++ {
		a := genElems(32)
		t.Log("Before:", a)
		f(a)
		t.Log("Sorted:", a)
		checkSorted(a, t)
	}
}

func TestSorting(t *testing.T) {
	t.Run("selection sort", func(t *testing.T) { test(selectionSort, t) })
	t.Run("insertion sort", func(t *testing.T) { test(insertionSort, t) })
	t.Run("shell sort", func(t *testing.T) { test(shellSort, t) })
	t.Run("shell sort 2", func(t *testing.T) { test(shellSort2, t) })
	t.Run("merge sort", func(t *testing.T) { test(mergesort, t) })
	t.Run("merge sort with buffer", func(t *testing.T) { test(mergesortBuf, t) })
	t.Run("merge sort with buffer from bottom", func(t *testing.T) { test(mergesortFromBottom, t) })
}

func checkSorted(a []int, t *testing.T) {
	n := len(a)
	for i := 1; i < n; i++ {
		e1 := a[i-1]
		e2 := a[i]
		b := e1 <= e2
		assert.Truef(t, b, "at index %d, elem %d is smaller than its previous elem %d", i, e2, e1)
		if !b {
			t.FailNow()
		}
	}
}

// genElems generates a new slice with n elements
func genElems(n int) []int {
	a := make([]int, n, n)
	max := n * 4
	for i := range a {
		a[i] = rand.Intn(max)
	}
	return a
}

func benchmark(f func([]int), n int, b *testing.B) {
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

func Benchmark_SelectionSort(b *testing.B) {
	f := selectionSort
	b.Run("100", func(b *testing.B) { benchmark(f, 100, b) })
	b.Run("1000", func(b *testing.B) { benchmark(f, 1000, b) })
	b.Run("10000", func(b *testing.B) { benchmark(f, 10000, b) })
}

func Benchmark_InsertionSort(b *testing.B) {
	f := insertionSort
	b.Run("100", func(b *testing.B) { benchmark(f, 100, b) })
	b.Run("1000", func(b *testing.B) { benchmark(f, 1000, b) })
	b.Run("10000", func(b *testing.B) { benchmark(f, 10000, b) })
}

func Benchmark_ShellSort(b *testing.B) {
	f := shellSort
	b.Run("100", func(b *testing.B) { benchmark(f, 100, b) })
	b.Run("1000", func(b *testing.B) { benchmark(f, 1000, b) })
	b.Run("10000", func(b *testing.B) { benchmark(f, 10000, b) })
}

func Benchmark_MergeSort(b *testing.B) {
	f := mergesort
	b.Run("100", func(b *testing.B) { benchmark(f, 100, b) })
	b.Run("1000", func(b *testing.B) { benchmark(f, 1000, b) })
	b.Run("10000", func(b *testing.B) { benchmark(f, 10000, b) })
}

func Benchmark_MergeSortWithBuffer(b *testing.B) {
	f := mergesortBuf
	b.Run("100", func(b *testing.B) { benchmark(f, 100, b) })
	b.Run("1000", func(b *testing.B) { benchmark(f, 1000, b) })
	b.Run("10000", func(b *testing.B) { benchmark(f, 10000, b) })
}

func Benchmark_MergeSortFromBottom(b *testing.B) {
	f := mergesortFromBottom
	b.Run("100", func(b *testing.B) { benchmark(f, 100, b) })
	b.Run("1000", func(b *testing.B) { benchmark(f, 1000, b) })
	b.Run("10000", func(b *testing.B) { benchmark(f, 10000, b) })
}
