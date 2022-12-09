package day9

import (
	"github.com/cuzfrog/go-playground/adventofcode/2022/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMove(t *testing.T) {
	k := newKnot()
	assert.Equal(t, shared.Coord{}, k.cur)

	k.move(Up)
	assert.Equal(t, shared.Coord{0, 1}, k.cur)

	k.move(Right)
	k.move(Right)
	assert.Equal(t, shared.Coord{2, 1}, k.cur)

	k.move(Down)
	k.move(Down)
	k.move(Down)
	assert.Equal(t, shared.Coord{2, -2}, k.cur)

	k.move(Left)
	assert.Equal(t, shared.Coord{1, -2}, k.cur)
}

func TestFollowRight(t *testing.T) {
	head, tail := newKnot(), newKnot()

	head.move(Right)
	tail.follow(head)
	assert.Equal(t, shared.Coord{}, tail.cur)

	head.move(Right)
	tail.follow(head)
	assert.Equal(t, shared.Coord{1, 0}, tail.cur)
}

func TestFollowLeft(t *testing.T) {
	head, tail := newKnot(), newKnot()

	head.move(Left)
	tail.follow(head)
	assert.Equal(t, shared.Coord{}, tail.cur)

	head.move(Left)
	tail.follow(head)
	assert.Equal(t, shared.Coord{-1, 0}, tail.cur)
}

func TestFollowUp(t *testing.T) {
	head, tail := newKnot(), newKnot()

	head.move(Up)
	tail.follow(head)
	assert.Equal(t, shared.Coord{}, tail.cur)

	head.move(Up)
	tail.follow(head)
	assert.Equal(t, shared.Coord{0, 1}, tail.cur)
}

func TestFollowDown(t *testing.T) {
	head, tail := newKnot(), newKnot()

	head.move(Down)
	tail.follow(head)
	assert.Equal(t, shared.Coord{}, tail.cur)

	head.move(Down)
	tail.follow(head)
	assert.Equal(t, shared.Coord{0, -1}, tail.cur)
}

func TestFollowRightUp(t *testing.T) {
	head, tail := newKnot(), newKnot()

	head.cur.X = 1
	head.cur.Y = 1
	tail.follow(head)
	assert.Equal(t, shared.Coord{}, tail.cur)

	head.cur.X = 2
	head.cur.Y = 1
	tail.follow(head)
	assert.Equal(t, shared.Coord{1, 1}, tail.cur)

	head.cur.X = 1
	head.cur.Y = 2
	tail = newKnot()
	tail.follow(head)
	assert.Equal(t, shared.Coord{1, 1}, tail.cur)
}

func TestFollowRightDown(t *testing.T) {
	head, tail := newKnot(), newKnot()

	head.cur.X = 1
	head.cur.Y = -1
	tail.follow(head)
	assert.Equal(t, shared.Coord{}, tail.cur)

	head.cur.X = 2
	head.cur.Y = -1
	tail.follow(head)
	assert.Equal(t, shared.Coord{1, -1}, tail.cur)

	head.cur.X = 1
	head.cur.Y = -2
	tail = newKnot()
	tail.follow(head)
	assert.Equal(t, shared.Coord{1, -1}, tail.cur)
}

func TestFollowLeftUp(t *testing.T) {
	head, tail := newKnot(), newKnot()

	head.cur.X = -1
	head.cur.Y = 1
	tail.follow(head)
	assert.Equal(t, shared.Coord{}, tail.cur)

	head.cur.X = -2
	head.cur.Y = 1
	tail.follow(head)
	assert.Equal(t, shared.Coord{-1, 1}, tail.cur)

	head.cur.X = -1
	head.cur.Y = 2
	tail = newKnot()
	tail.follow(head)
	assert.Equal(t, shared.Coord{-1, 1}, tail.cur)
}

func TestFollowLeftDown(t *testing.T) {
	head, tail := newKnot(), newKnot()

	head.cur.X = -1
	head.cur.Y = -1
	tail.follow(head)
	assert.Equal(t, shared.Coord{}, tail.cur)

	head.cur.X = -2
	head.cur.Y = -1
	tail.follow(head)
	assert.Equal(t, shared.Coord{-1, -1}, tail.cur)

	head.cur.X = -1
	head.cur.Y = -2
	tail = newKnot()
	tail.follow(head)
	assert.Equal(t, shared.Coord{-1, -1}, tail.cur)
}

func TestInput(t *testing.T) {
	steps := parseSteps("./test-input")
	_, tail := performSteps(steps)
	assert.Equal(t, 13, tail.countNoDuplicatePosition())
}

func TestSolution1(t *testing.T) {
	steps := parseSteps("./input")
	_, tail := performSteps(steps)
	println(tail.countNoDuplicatePosition())
}
