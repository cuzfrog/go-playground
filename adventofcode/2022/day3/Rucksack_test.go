package day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInput(t *testing.T) {
	items := findAllShared("./test-input")
	sum := sumPriority(items)
	assert.Equal(t, 157, sum)
}

func TestAllItemPriorities(t *testing.T) {
	items := findAllShared("./input")
	sum := sumPriority(items)
	println(sum)
}
