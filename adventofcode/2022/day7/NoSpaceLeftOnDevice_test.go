package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildFilesystem(t *testing.T) {
	root := buildFilesystem("./test-input")
	assert.Equal(t, 48381165, root.size)
}
