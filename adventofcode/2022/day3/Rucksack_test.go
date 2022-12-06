package day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInput1(t *testing.T) {
	items := findAllShared("./test-input")
	sum := sumPriority(items)
	assert.Equal(t, 157, sum)
}

func TestAllItemPriorities(t *testing.T) {
	items := findAllShared("./input")
	sum := sumPriority(items)
	println(sum)
}

func TestInput2(t *testing.T) {
	badges := findAllBadges("./test-input")
	sum := sumPriority(badges)
	assert.Equal(t, 70, sum)
}

func TestAllBadgePriorities(t *testing.T) {
	items := findAllBadges("./input")
	sum := sumPriority(items)
	println(sum)
}
