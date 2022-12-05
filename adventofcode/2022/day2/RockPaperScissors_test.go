package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateScores(t *testing.T) {
	rounds := parseRounds("./test-input")
	sum := 0
	for _, r := range rounds {
		sum += score(r)
	}
	assert.Equal(t, 15, sum)
}

func TestGetScores(t *testing.T) {
	rounds := parseRounds("./input")
	sum := 0
	for _, r := range rounds {
		sum += score(r)
	}
	println(sum)
}
