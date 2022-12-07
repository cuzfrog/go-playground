package day6

import (
	"github.com/cuzfrog/go-playground/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetectMarker(t *testing.T) {
	marker, cnt := detectMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4)
	assert.Equal(t, "jpqm", marker)
	assert.Equal(t, 7, cnt)

	marker, cnt = detectMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 4)
	assert.Equal(t, 5, cnt)

	marker, cnt = detectMarker("nppdvjthqldpwncqszvftbrmjlhg", 4)
	assert.Equal(t, 6, cnt)

	marker, cnt = detectMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4)
	assert.Equal(t, 10, cnt)
}

func TestSolution(t *testing.T) {
	lines := utils.LoadFileLines("./input")
	_, cnt := detectMarker(lines[0], 4)
	println(cnt)

	_, cnt = detectMarker(lines[0], 14)
	println(cnt)
}
