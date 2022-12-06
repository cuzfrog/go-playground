package day5

import (
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseStacks(t *testing.T) {
	ss := parseStacks("./test-input-stacks")
	assert.Equal(t, []uint8{'N', 'Z'}, utils.SliceFrom[uint8](ss[0]))
	assert.Equal(t, []uint8{'D', 'C', 'M'}, utils.SliceFrom[uint8](ss[1]))
	assert.Equal(t, []uint8{'P'}, utils.SliceFrom[uint8](ss[2]))
}

func TestParseMoves(t *testing.T) {
	ms := parseMoves("./test-input-moves")

	assert.Equal(t, 1, ms[0].cnt)
	assert.Equal(t, 2, ms[0].src)
	assert.Equal(t, 1, ms[0].tgt)

	assert.Equal(t, 1, ms[3].cnt)
	assert.Equal(t, 1, ms[3].src)
	assert.Equal(t, 2, ms[3].tgt)
}

func TestPerformMove(t *testing.T) {
	ss := parseStacks("./test-input-stacks")
	ss.performMove(move{2, 2, 1})
	assert.Equal(t, []uint8{'C', 'D', 'N', 'Z'}, utils.SliceFrom[uint8](ss[0]))
}

func TestInput1(t *testing.T) {
	ss := parseStacks("./test-input-stacks")
	ms := parseMoves("./test-input-moves")
	for _, mv := range ms {
		ss.performMove(mv)
	}
	assert.Equal(t, "CMZ", ss.peerTop())
}

func TestSolution(t *testing.T) {
	ss := parseStacks("./input-stacks")
	ms := parseMoves("./input-moves")
	for _, mv := range ms {
		ss.performMove(mv)
	}
	res := ss.peerTop()
	println(res)
	assert.Equal(t, "MQSHJMWNH", res)
}

func TestSolution9001(t *testing.T) {
	ss := parseStacks("./input-stacks")
	ms := parseMoves("./input-moves")
	for _, mv := range ms {
		ss.performMove9001(mv)
	}
	res := ss.peerTop()
	println(res)
	assert.Equal(t, "LLWJRBHVZ", res)
}
