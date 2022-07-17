package binary_tree_traversal

import (
	"github.com/cuzfrog/go-playground/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tree = &node{
	1,
	&node{
		2,
		&node{4, nil, nil},
		&node{5, nil, nil},
	},
	&node{3, nil, nil},
}

func Test_recursivePreorderTraverse(t *testing.T) {
	l := utils.NewList[int](5)
	recursivePreorderTraverse(tree, l)
	assert.Equal(t, []int{1, 2, 4, 5, 3}, l.Arr)
}

func Test_preorderTraverse(t *testing.T) {
	l := preorderTraverse(tree, 5)
	assert.Equal(t, []int{1, 2, 4, 5, 3}, l.Arr)
}
