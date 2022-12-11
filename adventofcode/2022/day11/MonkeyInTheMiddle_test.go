package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMonkeys(t *testing.T) {
	ms := parseMonkeys("./test-input")
	m0 := ms[0]
	assert.True(t, m0.items.Contains(79))
	assert.True(t, m0.items.Contains(98))
	assert.Equal(t, 2, m0.items.Size())

	assert.Equal(t, 19, m0.ops(1))
	assert.Equal(t, 23, m0.testDividant)
	assert.Equal(t, 2, m0.tgtT)
	assert.Equal(t, 3, m0.tgtF)

	m2 := ms[2]
	assert.Equal(t, 9, m2.ops(3))

	m3 := ms[3]
	assert.True(t, m3.items.Contains(74))
	assert.Equal(t, 1, m3.items.Size())

	assert.Equal(t, 4, m3.ops(1))
	assert.Equal(t, 17, m3.testDividant)
	assert.Equal(t, 0, m3.tgtT)
	assert.Equal(t, 1, m3.tgtF)
}
