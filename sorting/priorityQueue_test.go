package sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_reheapifyLast(t *testing.T) {
	h := []int{0, 10, 5, 7, 3}
	h = append(h, 2)
	reheapifyLast(h)
	assert.Equal(t, )
}
