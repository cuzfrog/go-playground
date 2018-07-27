package unionfind

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math/rand"
)

func createInstances(capacity int) []unions {
	return []unions{
		newQfUnions(capacity),
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
		Assert.Equal(96, u.count(), "Component num is wrong")
		Assert.True(u.connected(67, 6))
		Assert.False(u.connected(8, 9))
	}
}

func benchmark(uCtor func(int) unions, cap int, b *testing.B) {
	u := uCtor(cap)
	pairs := genPairs(cap)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pIdx := i % cap
		if pIdx == 0 {
			u.reset()
		}
		b.StartTimer()
		u.union(pIdx, pairs[pIdx])
		u.connected(pIdx, pairs[pIdx])
		u.count()
	}
}

func genPairs(n int) []int {
	pairs := make([]int, n)
	for i := 0; i < n; i ++ {
		pairs[i] = rand.Intn(n)
	}
	return pairs
}

func BenchmarkQf(b *testing.B) {
	b.Run("100", func(b *testing.B) {benchmark(newQfUnions, 100, b)})
	b.Run("1000", func(b *testing.B) {benchmark(newQfUnions, 1000, b)})
	b.Run("10000", func(b *testing.B) {benchmark(newQfUnions, 10000, b)})
	b.Run("100000", func(b *testing.B) {benchmark(newQfUnions, 100000, b)})
}