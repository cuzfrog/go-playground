package day10

import (
	"github.com/cuzfrog/go-playground/utils"
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

func TestInput2(t *testing.T) {
	inss := parseInstructions("./test-input")
	sprites := parseSprites(inss)
	f := scan(sprites)
	output := f.Sprint()
	expected := utils.LoadFileContent("./test-output")
	assert.Equal(t, expected, output)
}

func TestSolution2(t *testing.T) {
	inss := parseInstructions("./input")
	sprites := parseSprites(inss)
	f := scan(sprites)
	println(f.Sprint())
}
