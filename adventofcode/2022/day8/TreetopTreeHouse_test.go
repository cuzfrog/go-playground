package day8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseForest(t *testing.T) {
	f := parseForest("./test-input")
	assert.Equal(t, int8(5), f.trees[1][2])
	assert.Equal(t, int8(9), f.trees[3][4])
}

func TestScan(t *testing.T) {
	f := parseForest("./test-input")
	cnt := scan(f)
	assert.Equal(t, 21, cnt)
}

func TestSolution1(t *testing.T) {
	f := parseForest("./input")
	cnt := scan(f)
	println(cnt)
}
