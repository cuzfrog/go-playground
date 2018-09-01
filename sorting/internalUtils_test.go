package sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_swim(t *testing.T) {
	h := []int{0}
	h = append(h, 5)
	swimMaxLast(h)
	assert.Equal(t, []int{0, 5}, h)

	h = append(h, 10)
	swimMaxLast(h)
	assert.Equal(t, []int{0, 10, 5}, h)

	h = append(h, 7)
	swimMaxLast(h)
	assert.Equal(t, []int{0, 10, 5, 7}, h)

	h = append(h, 3)
	swimMaxLast(h)
	assert.Equal(t, []int{0, 10, 5, 7, 3}, h)

	h = append(h, 2)
	swimMaxLast(h)
	assert.Equal(t, []int{0, 10, 5, 7, 3, 2}, h)

	h = append(h, 8)
	swimMaxLast(h)
	assert.Equal(t, []int{0, 10, 5, 8, 3, 2, 7}, h)

	h = append(h, 15)
	i := swimMaxLast(h)
	assert.Equal(t, []int{0, 15, 5, 10, 3, 2, 7, 8}, h)
	assert.Equal(t, 1, i)
}

func Test_sink(t *testing.T) {
	h := []int{0, 2, 5, 10, 3}
	v, i := sinkMax(h, 1, nil)
	assert.Equal(t, []int{0, 10, 5, 2, 3}, h)
	assert.Equal(t, 2, v)
	assert.Equal(t, 3, i)

	sinkMax(h, 2, nil)
	assert.Equal(t, []int{0, 10, 5, 2, 3}, h)
}
