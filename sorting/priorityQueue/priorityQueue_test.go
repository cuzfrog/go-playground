package priorityQueue

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_swimLast(t *testing.T) {
	h := []int{0}
	h = append(h, 5)
	swimLast(h)
	assert.Equal(t, []int{0, 5}, h)

	h = append(h, 10)
	swimLast(h)
	assert.Equal(t, []int{0, 10, 5}, h)

	h = append(h, 7)
	swimLast(h)
	assert.Equal(t, []int{0, 10, 5, 7}, h)

	h = append(h, 3)
	swimLast(h)
	assert.Equal(t, []int{0, 10, 5, 7, 3}, h)

	h = append(h, 2)
	swimLast(h)
	assert.Equal(t, []int{0, 10, 5, 7, 3, 2}, h)

	h = append(h, 8)
	swimLast(h)
	assert.Equal(t, []int{0, 10, 5, 8, 3, 2, 7}, h)

	h = append(h, 15)
	swimLast(h)
	assert.Equal(t, []int{0, 15, 5, 10, 3, 2, 7, 8}, h)
}

func Test_sink(t *testing.T) {
	h := []int{0, 2, 5, 10, 3}
	sink(h, 1)
	assert.Equal(t, []int{0, 10, 5, 2, 3}, h)

	sink(h, 2)
	assert.Equal(t, []int{0, 10, 5, 2, 3}, h)
}

func Test_priorityQueue(t *testing.T){

}