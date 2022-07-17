package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasUniqueChar(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"abc", true},
		{"abcb", false},
	}
	for _, tt := range tests {
		assert.Equal(t, hasUniqueChar(&tt.s), tt.want)
		assert.Equal(t, hasUniqueChar2(&tt.s), tt.want)
	}
}
