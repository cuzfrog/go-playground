package sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"math/rand"
)

func Test_swimLast(t *testing.T) {
	h := []int{0}
	h = append(h, 5)
	swimLast(h)
	assert.Equal(t, []int{0, 5}, h)

	h = append(h, 10)
	swimLast(h)
	assert.Equal(t, []int{0, 10, 5}, h)

	h = append(h, 7)
	swimLast(h)
	assert.Equal(t, []int{0, 10, 5, 7}, h)

	h = append(h, 3)
	swimLast(h)
	assert.Equal(t, []int{0, 10, 5, 7, 3}, h)

	h = append(h, 2)
	swimLast(h)
	assert.Equal(t, []int{0, 10, 5, 7, 3, 2}, h)

	h = append(h, 8)
	swimLast(h)
	assert.Equal(t, []int{0, 10, 5, 8, 3, 2, 7}, h)

	h = append(h, 15)
	swimLast(h)
	assert.Equal(t, []int{0, 15, 5, 10, 3, 2, 7, 8}, h)
}

func Test_sink(t *testing.T) {
	h := []int{0, 2, 5, 10, 3}
	sink(h, 1)
	assert.Equal(t, []int{0, 10, 5, 2, 3}, h)

	sink(h, 2)
	assert.Equal(t, []int{0, 10, 5, 2, 3}, h)
}

func Test_priorityQueue(t *testing.T) {
	for n := 0; n < 100; n++ {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			t.Parallel()
			a := GenElems(rand.Intn(32))
			n := len(a)
			pq := New(a)
			assert.Equal(t, n, pq.size())
			assertHeapProperty(pq.heap, t)

			quicksort(a)
			max := a[n-1]
			assert.Equal(t, max, pq.popMax())
			assert.Equal(t, n-1, pq.size())
		})
	}
}

func assertHeapProperty(h []int, t *testing.T) {
	n := len(h)
	for i := 1; i <= n/2; i++ {
		msg := "Elem %v(%v) is less than elem %v(%v) in heap %v"
		if i*2 < n {
			assert.Truef(t, h[i] >= h[i*2], msg, i, h[i], i*2, h[i*2], h)
		}
		if i*2+1 < n {
			assert.Truef(t, h[i] >= h[i*2+1], msg, i, h[i], i*2+1, h[i*2+1], h)
		}
	}
}
