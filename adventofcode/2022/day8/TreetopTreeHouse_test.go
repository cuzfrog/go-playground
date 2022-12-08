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

func TestScanVisible(t *testing.T) {
	f := parseForest("./test-input")
	cnt := scanVisible(f)
	assert.Equal(t, 21, cnt)
}

func TestSolution1(t *testing.T) {
	f := parseForest("./input")
	cnt := scanVisible(f)
	println(cnt)
	assert.Equal(t, 1708, cnt)
}

func TestScenicScore(t *testing.T) {
	f := parseForest("./test-input")
	s := f.scenicScore(1, 2)
	assert.Equal(t, 4, s)
	s = f.scenicScore(3, 2)
	assert.Equal(t, 8, s)
}

func TestSolution2(t *testing.T) {
	f := parseForest("./input")
	max := 0
	for i := 0; i < f.rowCnt; i++ {
		for j := 0; j < f.colCnt; j++ {
			s := f.scenicScore(i, j)
			if s > max {
				max = s
			}
		}
	}
	println(max)
}
