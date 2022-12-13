package day13

import (
	"github.com/cuzfrog/tgods/collections"
	"github.com/cuzfrog/tgods/transform"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSignal(t *testing.T) {
	s := parseSignal([]byte("[[1],[2,3,4],5,6]"), span{0, 17})
	assert.True(t, s.isList)
	assert.Equal(t, 4, len(s.values))
	assert.True(t, s.values[0].isList)
	assert.Equal(t, 1, s.values[0].values[0].v)
	assert.Equal(t, 2, s.values[1].values[0].v)
	assert.Equal(t, 4, s.values[1].values[2].v)
	assert.False(t, s.values[2].isList)
	assert.Equal(t, 5, s.values[2].v)
}

func TestSplitElems(t *testing.T) {
	sps := splitElems([]byte("[[1],[2,3,4],5,6]"), span{6, 11})
	assert.Equal(t, span{6, 7}, sps[0])
	assert.Equal(t, span{8, 9}, sps[1])
	assert.Equal(t, span{10, 11}, sps[2])

	sps = splitElems([]byte("[[1],[2,3,4],5,6]"), span{1, 16})
	assert.Equal(t, 4, len(sps))
	assert.Equal(t, span{5, 12}, sps[1])
}

func newPair(l, r string) pair {
	return pair{newSignal(l), newSignal(r)}
}

func newSignal(s string) *signal {
	return parseSignal([]byte(s), span{0, len(s)})
}

func TestIsCorrect(t *testing.T) {
	assert.Equal(t, 0, findOrder(newPair("[[1],[2,3,4],5,6]", "[[1],[2,3,4],5,6]")))
	assert.Equal(t, -1, findOrder(newPair("[[1]]", "[[2]]")))
	assert.Equal(t, 1, findOrder(newPair("[[3]]", "[[2]]")))
	assert.Equal(t, -1, findOrder(newPair("[]", "[[]]")))
	assert.Equal(t, -1, findOrder(newPair("[1,2]", "[3]")))

	assert.Equal(t, -1, findOrder(newPair("[[1],[2,3,4]]", "[[1],4]")))
	assert.Equal(t, -1, findOrder(newPair("[[1],[2,3,4],5,6]", "[[1],[2,3,4],7]")))
	assert.Equal(t, 1, findOrder(newPair("[9]", "[[8,7,6]]")))
	assert.Equal(t, -1, findOrder(newPair("[[4,4],4,4]", "[[4,4],4,4,4]")))
	assert.Equal(t, -1, findOrder(newPair("[]", "[3]")))
	assert.Equal(t, 1, findOrder(newPair("[[[]]]", "[[]]")))
}

func TestInput1(t *testing.T) {
	ps := parseSignals("./test-input")
	rightIndices := collections.NewArrayListOf[int]()
	for i, p := range ps {
		if findOrder(p) <= 0 {
			rightIndices.Add(i + 1)
		}
	}
	assert.Equal(t, []int{1, 2, 4, 6}, utils.SliceFrom[int](rightIndices))

	ps = parseSignals("./input")
	rightIndices.Clear()
	for i, p := range ps {
		if findOrder(p) <= 0 {
			rightIndices.Add(i + 1)
		}
	}
	sum := transform.Reduce[int, int](rightIndices, 0, func(a, i int) int { return a + i })
	println(sum)
}

var d1 = newSignal("[[2]]")
var d2 = newSignal("[[6]]")

func TestInput2(t *testing.T) {
	ps := parseSignals("./test-input")
	signals := collections.NewArrayListOfSize[*signal](len(ps) + 2)
	transform.FlatMapSliceTo[pair, *signal](ps, signals, pairToSignals)
	signals.Add(d1)
	signals.Add(d2)
	signals.Sort(signalLess)

	ss := utils.SliceFrom[*signal](signals)
	assert.Equal(t, d1, ss[9])
	assert.Equal(t, d2, ss[13])
}

func TestSolution2(t *testing.T) {
	ps := parseSignals("./input")
	signals := collections.NewArrayListOfSize[*signal](len(ps) + 2)
	transform.FlatMapSliceTo[pair, *signal](ps, signals, pairToSignals)

	signals.Add(d1)
	signals.Add(d2)
	signals.Sort(signalLess)

	key := 1
	signals.Each(func(i int, s *signal) {
		if s == d1 || s == d2 {
			key *= (i + 1)
		}
	})
	println(key)
	assert.Equal(t, 20952, key)
}
