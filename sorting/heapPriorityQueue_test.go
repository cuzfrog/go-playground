package sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"math/rand"
)

func Test_heapPriorityQueue(t *testing.T) {
	for m := 0; m < 100; m++ {
		t.Run(fmt.Sprint(m), func(t *testing.T) {
			t.Parallel()
			a := GenElems(rand.Intn(32))
			n := len(a)
			pq := NewHeapPriorityQueue(a)
			assert.Equal(t, n, pq.size())
			assertHeapProperty(pq.heap, t)

			if n > 0 {
				quicksort(a)
				max := a[n-1]
				assert.Equal(t, max, pq.popMax())
				assert.Equal(t, n-1, pq.size())
			}
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
