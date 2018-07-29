package HOF

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFoldLeft(t *testing.T) {
	s := []int{1,5,8,4,3}
	sum := foldLeft(s, 0, func(prevId interface{}, next interface{}) interface{} {
		return prevId.(int) + next.(int)
	})
	assert.Equal(t, sum, 21)
}