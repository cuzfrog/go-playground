package day7

import (
	"fmt"
	"github.com/cuzfrog/tgods/transform"
	"github.com/stretchr/testify/assert"
	"testing"
)

func sumSize(acc int, next *file) int {
	return acc + next.size
}

func filterDirBelow100000(f *file) bool {
	return f.size <= 100000 && f.isDir
}

func TestBuildFilesystem(t *testing.T) {
	root := buildFilesystem("./test-input")
	assert.Equal(t, 48381165, root.size)
}

func TestTraverse(t *testing.T) {
	root := buildFilesystem("./test-input")
	files := traverseFilesystem(root, filterDirBelow100000)
	assert.Equal(t, 2, files.Size())
	totalSize := transform.Reduce[*file, int](files, 0, sumSize)
	assert.Equal(t, 95437, totalSize)
}

func TestSolution1(t *testing.T) {
	root := buildFilesystem("./input")
	files := traverseFilesystem(root, filterDirBelow100000)
	totalSize := transform.Reduce[*file, int](files, 0, sumSize)
	println(totalSize)
	assert.Equal(t, 1989474, totalSize)
}

func TestSolution2(t *testing.T) {
	root := buildFilesystem("./input")
	spaceToDel := 30000000 - (70000000 - root.size)
	println(fmt.Sprintf("need to at least delete space: %d", spaceToDel))

	files := traverseFilesystem(root, func(f *file) bool { return f.size >= spaceToDel && f.isDir })

	files.Sort(func(f1, f2 *file) bool { return f1.size < f2.size })
	smallest, _ := files.Head()
	println(smallest.size)
}
