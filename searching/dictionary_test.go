package searching

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"go-playground/sorting"
)

func Test_dictionary_implementations(t *testing.T) {
	t.Run("23tree", func(t *testing.T) { test(t, new23TreeDic) })
}

func test(t *testing.T, ctor func() dictionary) {
	asert := assert.New(t)
	d := ctor()
	asert.Equal(0, d.size())
	asert.False(d.contains(1))

	keys := sorting.Dedup(sorting.GenElems(100))
	size := len(keys)
	vals := make([]string, size)
	for i := range vals {
		k := keys[i]
		v := string(k)
		vals[i] = v
		d.put(k, v)
	}
	for i := 0; i < size/2; i++ {
		vals[i] = vals[i] + "a"
		d.put(keys[i], vals[i])
	}
	asert.Equal(size, d.size())
	for i := range keys {
		asert.True(d.contains(keys[i]))
		asert.Equal(vals[i], d.get(keys[i]))
	}

	for i := range keys {
		d.remove(keys[i])
	}
	asert.Equal(0, d.size())
}

func benchmark(b *testing.B, n int, ctor func() dictionary) {
	d := ctor()
	keys := sorting.GenNonDupElems(n)
	for i := range keys {
		d.put(keys[i], i)
	}
	testNums := sorting.GenElems(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		num := testNums[i]
		d.put(num, i)
		d.remove(num)
	}
}

func Benchmark_23treeDic(b *testing.B)  {
	ctor := new23TreeDic
	b.Run("100", func(b *testing.B) {benchmark(b, 100, ctor)})
	b.Run("1000", func(b *testing.B) {benchmark(b, 100, ctor)})
	b.Run("10000", func(b *testing.B) {benchmark(b, 100, ctor)})
}
