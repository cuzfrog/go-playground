package sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_reheapifyLast(t *testing.T) {
	h := []int{0}
	h = append(h, 5)
	reheapifyLast(h)
	assert.Equal(t, []int{0, 5}, h)

	h = append(h, 10)
	reheapifyLast(h)
	assert.Equal(t, []int{0, 10, 5}, h)

	h = append(h, 7)
	reheapifyLast(h)
	assert.Equal(t, []int{0, 10, 5, 7}, h)

	h = append(h, 3)
	reheapifyLast(h)
	assert.Equal(t, []int{0, 10, 5, 7, 3}, h)

	h = append(h, 2)
	reheapifyLast(h)
	assert.Equal(t, []int{0, 10, 5, 7, 3, 2}, h)

	h = append(h, 8)
	reheapifyLast(h)
	assert.Equal(t, []int{0, 10, 5, 8, 3, 2, 7}, h)

	h = append(h, 15)
	reheapifyLast(h)
	assert.Equal(t, []int{0, 15, 5, 10, 3, 2, 7, 8}, h)
}
