package day7

import (
	"github.com/cuzfrog/go-playground/utils"
	"github.com/cuzfrog/tgods/collections"
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
)

type file struct {
	path   string
	isDir  bool
	size   int
	parent *file
	files  types.Set[*file]
}

func fileHash(f *file) uint {
	return funcs.NewStrHash()(f.path)
}
func fileEqual(a, b *file) bool {
	return funcs.ValueEqual(a.path, b.path)
}

func traverseFilesystem(root *file, filterFn func(f *file) bool) types.ArrayList[*file] {
	files := collections.NewArrayListOfEq(0, fileEqual)
	if filterFn(root) {
		files.Add(root)
	}
	cur := root
	queue := collections.NewArrayListQueueOfEq(0, fileEqual)
	for queue.Size() > 0 || (cur != nil && cur.files != nil) {
		it := cur.files.Iterator()
		for it.Next() {
			f := it.Value()
			if f.isDir {
				queue.Enqueue(f)
			}
			if filterFn(f) {
				files.Add(f)
			}
		}
		cur, _ = queue.Dequeue()
	}
	return files
}

func buildFilesystem(path string) *file {
	lines := utils.LoadFileLines(path)
	l := len(lines) - 1

	root := &file{"/", true, 0, nil, collections.NewHashSet[*file](fileHash, fileEqual)}
	cur := root
	for i := 0; i < l; i++ {
		line := lines[i]
		if line[0] == '$' {
			cmd := line[2:4]
			if cmd == "cd" {
				cur = navigate(root, cur, line[5:])
			} else if cmd != "ls" {
				panic("unknown cmd")
			}
		} else {
			if line[:3] != "dir" {
				sizeStr, name := utils.SplitString2(line, " ")
				addFile(cur, name, utils.StrToInt(sizeStr))
			}
		}
	}

	return root
}

func addFile(cur *file, name string, size int) {
	f := &file{cur.path + "/" + name, false, size, cur, nil}
	if !cur.files.Contains(f) {
		cur.files.Add(f)
		cur.size += size
		p := cur.parent
		for p != nil {
			p.size += size
			p = p.parent
		}
	}
}

func navigate(root, cur *file, cmdV string) *file {
	if cmdV == "/" {
		cur = root
	} else if cmdV == ".." {
		cur = cur.parent
	} else {
		dir := &file{cur.path + "/" + cmdV, true, 0, cur, collections.NewHashSet[*file](fileHash, fileEqual)}
		cur.files.Add(dir)
		cur = dir
	}
	return cur
}
