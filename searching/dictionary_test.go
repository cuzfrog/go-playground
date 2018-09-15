package searching

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"go-playground/sorting"
)

func Test_dictionary_implementations(t *testing.T) {
	t.Run("23tree", func(t *testing.T) { testDictionary(t, new23TreeDic) })
}

func testDictionary(t *testing.T, ctor func() dictionary) {
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

	for i := range keys{
		d.remove(keys[i])
	}
	asert.Equal(0, d.size())
}
