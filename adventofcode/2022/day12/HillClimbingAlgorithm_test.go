package day12

import (
	"github.com/cuzfrog/go-playground/adventofcode/2022/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadMap(t *testing.T) {
	m := loadMap("./test-input")
	assert.Equal(t, 5, m.rowCnt)
	assert.Equal(t, 8, m.colCnt)
	assert.Equal(t, shared.Coord{0, 0}, m.start)
	assert.Equal(t, shared.Coord{5, 2}, m.end)

	assert.Equal(t, int8('s'), m.height(shared.Coord{3, 2}))
	assert.Equal(t, int8('z'), m.height(shared.Coord{4, 2}))
}

func TestMoveUp(t *testing.T) {
	m := loadMap("./test-input")
	n, ok := moveUp(m, shared.Coord{2, 2})
	assert.True(t, ok)
	assert.Equal(t, shared.Coord{2, 1}, n)

	n, ok = moveUp(m, shared.Coord{4, 3})
	assert.False(t, ok)
}

func TestMoveRight(t *testing.T) {
	m := loadMap("./test-input")
	n, ok := moveRight(m, shared.Coord{0, 1})
	assert.True(t, ok)
	assert.Equal(t, shared.Coord{1, 1}, n)

	n, ok = moveRight(m, shared.Coord{0, 2})
	assert.False(t, ok)
}

func TestMoveDown(t *testing.T) {
	m := loadMap("./test-input")
	n, ok := moveDown(m, shared.Coord{1, 0})
	assert.True(t, ok)
	assert.Equal(t, shared.Coord{1, 1}, n)

	n, ok = moveDown(m, shared.Coord{4, 0})
	assert.False(t, ok)
}

func TestMoveLeft(t *testing.T) {
	m := loadMap("./test-input")
	n, ok := moveLeft(m, shared.Coord{4, 0})
	assert.True(t, ok)
	assert.Equal(t, shared.Coord{3, 0}, n)

	n, ok = moveLeft(m, shared.Coord{7, 1})
	assert.False(t, ok)
}

func TestInput1(t *testing.T) {
	m := loadMap("./test-input")
	cnt, endStep := moveCnt(m)
	println(endStep.Route())
	assert.Equal(t, 31, cnt)
}

func TestSolution1(t *testing.T) {
	m := loadMap("./input")
	cnt, _ := moveCnt(m)
	println(cnt)
	assert.Equal(t, 517, cnt)
}
