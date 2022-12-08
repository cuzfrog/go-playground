package day8

import (
	"fmt"
	"github.com/cuzfrog/go-playground/utils"
	"github.com/cuzfrog/tgods/collections"
	"github.com/cuzfrog/tgods/types"
)

type forest struct {
	trees  [][]int8
	rowCnt int
	colCnt int
}

type coord struct {
	x int
	y int
}

func (c coord) Hash() uint {
	return uint(c.x*17 + c.y)
}
func (c coord) Equal(o coord) bool {
	return c.x == o.x && c.y == o.y
}
func (c coord) String() string {
	return fmt.Sprintf("%d.%d", c.x, c.y)
}

func (f *forest) scenicScore(r, c int) int {
	h := f.trees[r][c]

	downScore := 0
	for i := r + 1; i < f.rowCnt; i++ {
		nh := f.trees[i][c]
		if h > nh {
			downScore++
		} else if h <= nh {
			downScore++
			break
		}
	}

	upScore := 0
	for i := r - 1; i >= 0; i-- {
		nh := f.trees[i][c]
		if h > nh {
			upScore++
		} else if h <= nh {
			upScore++
			break
		}
	}

	rightScore := 0
	for j := c + 1; j < f.colCnt; j++ {
		nh := f.trees[r][j]
		if h > nh {
			rightScore++
		} else if h <= nh {
			rightScore++
			break
		}
	}

	leftScore := 0
	for j := c - 1; j >= 0; j-- {
		nh := f.trees[r][j]
		if h > nh {
			leftScore++
		} else if h <= nh {
			leftScore++
			break
		}
	}

	return downScore * upScore * rightScore * leftScore
}

func scanVisible(f forest) int {
	visible := collections.NewHashSetC[coord]()

	bar := int8(-1)
	for i := 0; i < f.rowCnt; i++ {
		bar = -1
		for j := 0; j < f.colCnt; j++ {
			h := f.trees[i][j]
			if h > bar {
				visible.Add(coord{i, j})
				bar = h
			}
		}

		bar = -1
		for j := f.colCnt - 1; j >= 0; j-- {
			h := f.trees[i][j]
			if h > bar {
				visible.Add(coord{i, j})
				bar = h
			}
		}
	}
	for j := 0; j < f.colCnt; j++ {
		bar = -1
		for i := 0; i < f.rowCnt; i++ {
			h := f.trees[i][j]
			if h > bar {
				visible.Add(coord{i, j})
				bar = h
			}
		}

		bar = -1
		for i := f.rowCnt - 1; i >= 0; i-- {
			h := f.trees[i][j]
			if h > bar {
				visible.Add(coord{i, j})
				bar = h
			}
		}
	}
	printVisible(f, visible)
	return visible.Size()
}

func printVisible(f forest, visible types.Set[coord]) {
	for i := 0; i < f.rowCnt; i++ {
		for j := 0; j < f.colCnt; j++ {
			if visible.Contains(coord{i, j}) {
				print(f.trees[i][j])
			} else {
				print(" ")
			}
		}
		println()
	}
}

func parseForest(path string) forest {
	lines := utils.LoadFileLines(path)
	l := len(lines) - 1
	colCnt := len(lines[0])
	f := make([][]int8, l)
	for i := 0; i < l; i++ {
		line := lines[i]
		row := make([]int8, colCnt)
		for j := 0; j < colCnt; j++ {
			row[j] = int8(line[j] - '0')
		}
		f[i] = row
	}
	return forest{f, l, colCnt}
}
