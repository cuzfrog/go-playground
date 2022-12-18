package day15

import (
	"github.com/cuzfrog/go-playground/adventofcode/2022/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parsePairs(t *testing.T) {
	ch, pairs := parseChart("./test-input")
	assert.Equal(t, 14, len(pairs))
	assert.Equal(t, shared.Coord{2, 18}, pairs[0].se)
	assert.Equal(t, shared.Coord{-2, 15}, pairs[0].be)
	assert.Equal(t, shared.Coord{14, 3}, pairs[12].se)
	assert.Equal(t, shared.Coord{15, 3}, pairs[12].be)

	assert.Equal(t, sensor, ch.Get(2, 0))
	assert.Equal(t, beacon, ch.Get(-2, 15))
	assert.Equal(t, beacon, ch.Get(25, 17))
	assert.Equal(t, empty, ch.Get(25, 22))
	assert.Equal(t, beacon, ch.Get(21, 22))
}
