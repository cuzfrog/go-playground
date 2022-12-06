package day4

import (
	"github.com/cuzfrog/tgods/transform"
	"github.com/stretchr/testify/assert"
	"testing"
)

func isPairCovered(p pair) bool {
	return p.coverTheOther()
}

func isPairOverlap(p pair) bool {
	return p.overlap()
}

func TestInput1(t *testing.T) {
	pairs := parsePairs("./test-input")
	cnt := transform.CountSlice(pairs, isPairCovered)
	assert.Equal(t, 2, cnt)
}

func TestInput2(t *testing.T) {
	pairs := parsePairs("./test-input")
	cnt := transform.CountSlice(pairs, isPairOverlap)
	assert.Equal(t, 4, cnt)
}

func TestCountCovers(t *testing.T) {
	pairs := parsePairs("./input")
	cnt := transform.CountSlice(pairs, isPairCovered)
	println(cnt)
}

func TestCountOverlap(t *testing.T) {
	pairs := parsePairs("./input")
	cnt := transform.CountSlice(pairs, isPairOverlap)
	println(cnt)
}
