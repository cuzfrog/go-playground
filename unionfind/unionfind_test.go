package unionfind

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func createInstances(n int) []unions {
	return []unions{
		newQfUnions(n),
		newQfUnions(n),
	}
}

func TestUnions(t *testing.T) {
	Assert := assert.New(t)
	for _, u := range createInstances(100) {
		u.union(3, 6)
		u.union(8, 11)
		u.union(3, 9)
		u.union(67, 9)
		u.union(6, 9)
		Assert.Equal(96, u.count())
		Assert.True(u.connected(67, 6))
		Assert.False(u.connected(8, 9))
	}
}

func benchmark(uCtor func(int) unions, cap int, b *testing.B) {
	var u unions
	pairs := genPairs(cap)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pIdx := i % cap
		if pIdx == 0 {
			u = uCtor(cap)
		}
		pair := pairs[pIdx]
		b.StartTimer()

		u.union(pair.p, pair.q)
		u.count()
	}
}

type pair struct {
	p int
	q int
}

func genPairs(n int) []pair {
	pairs := make([]pair, n)
	for i := 0; i < n; i++ {
		pairs[i] = pair{rand.Intn(n), rand.Intn(n)}
	}
	return pairs
}

func BenchmarkQf(b *testing.B) {
	f := newQfUnions
	b.Run("1000", func(b *testing.B) { benchmark(f, 1000, b) })
	b.Run("10000", func(b *testing.B) { benchmark(f, 10000, b) })
	b.Run("100000", func(b *testing.B) { benchmark(f, 100000, b) })
}

func BenchmarkQu(b *testing.B) {
	f := newQuUnions
	b.Run("1000", func(b *testing.B) { benchmark(f, 1000, b) })
	b.Run("10000", func(b *testing.B) { benchmark(f, 10000, b) })
	b.Run("100000", func(b *testing.B) { benchmark(f, 100000, b) })
}
