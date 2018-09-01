package sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_heapSort_sink(t *testing.T) {
	a := []int{3, 5, 8, 2, 5, 9}
	sink(a, 3)
	assert.Equal(t, []int{3, 5, 9, 2, 5, 8}, a)
	sink(a, 2)
	assert.Equal(t, []int{3, 5, 9, 2, 5, 8}, a)
	sink(a, 1)
	assert.Equal(t, []int{9, 5, 8, 2, 5, 3}, a)
}
