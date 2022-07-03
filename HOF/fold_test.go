package HOF

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoldLeft(t *testing.T) {
	s := []int{1, 5, 8, 4, 3}
	sum := foldLeft(s, 0, func(prevId interface{}, next interface{}) interface{} {
		return prevId.(int) + next.(int)
	})
	assert.Equal(t, sum, 21)

	sum2 := foldLeftGeneric(s, 0, func(prev int, next int) int {
		return (prev + next)
	})
	assert.Equal(t, sum2, 21)
}
