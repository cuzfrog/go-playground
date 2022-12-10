package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInput(t *testing.T) {
	inss := parseInstructions("./test-input")
	st := sumSignalStrength(inss)
	assert.Equal(t, 13140, st)
}

func TestSolution1(t *testing.T) {
	inss := parseInstructions("./input")
	st := sumSignalStrength(inss)
	println(st)
	assert.Equal(t, 12540, st)
}
