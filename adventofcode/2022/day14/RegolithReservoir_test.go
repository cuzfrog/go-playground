package day14

import (
	"fmt"
	"github.com/cuzfrog/go-playground/adventofcode/2022/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLines(t *testing.T) {
	lines, rec := parseLines("./test-input")
	l1 := lines[0]
	assert.Equal(t, shared.ParseCoord("498,4"), l1.v1)
	assert.Equal(t, shared.ParseCoord("498,6"), l1.v2)
	assert.True(t, l1.isVertical)

	l4 := lines[3]
	assert.Equal(t, shared.ParseCoord("502,4"), l4.v1)
	assert.Equal(t, shared.ParseCoord("502,9"), l4.v2)
	assert.True(t, l4.isVertical)

	assert.False(t, lines[1].isVertical)

	assert.Equal(t, 5, len(lines))

	assert.Equal(t, shared.ParseCoord("494,4"), rec.Ori)
	assert.Equal(t, 10, rec.Width)
	assert.Equal(t, 6, rec.Height)
}

func TestBuildCaveSlice(t *testing.T) {

	cs := buildCaveSlice(parseLines("./test-input"))
	println(fmt.Sprintf("%v", cs))
	assert.Equal(t, 10, cs.width)
	assert.Equal(t, 10, cs.height)

	assert.Equal(t, rock, cs.s[6][2])
	assert.Equal(t, rock, cs.s[9][0])
	assert.Equal(t, rock, cs.s[9][8])
	assert.Equal(t, air, cs.s[9][9])
	assert.Equal(t, air, cs.s[4][3])
	assert.Equal(t, rock, cs.s[4][4])
	assert.Equal(t, air, cs.s[5][6])
	assert.Equal(t, shared.ParseCoord("6,0"), cs.src)
}
