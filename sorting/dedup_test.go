package sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_dedup(t *testing.T) {
	a := []int{4, 7, 3, 6, 4, 3, 7, 5}
	d := Dedup(a)
	assert.Equal(t, []int{3, 4, 5, 6, 7}, d)
}
