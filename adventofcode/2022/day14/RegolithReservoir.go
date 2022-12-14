package day14

import (
	"github.com/cuzfrog/go-playground/adventofcode/2022/shared"
	"github.com/cuzfrog/go-playground/utils"
	"github.com/cuzfrog/tgods/collections"
	utils2 "github.com/cuzfrog/tgods/utils"
	"math"
	"strings"
)

const (
	air byte = iota
	rock
	sand
)

type caveSlice struct {
	s      [][]byte
	height int
	width  int
	src    shared.Coord
}

type line struct {
	v1         shared.Coord
	v2         shared.Coord
	isVertical bool
}

func buildCaveSlice(lines []line, rec shared.Rectangle) *caveSlice {
	height := rec.Height + rec.Ori.Y
	s := make([][]byte, height)
	for i := 0; i < height; i++ {
		s[i] = make([]byte, rec.Width)
	}

	for _, l := range lines {
		if l.isVertical {
			j := l.v1.X - rec.Ori.X
			y1, y2 := utils.MinMaxOf(l.v1.Y, l.v2.Y)
			for i := y1; i <= y2; i++ {
				s[i][j] = rock
			}
		} else {
			x1, x2 := utils.MinMaxOf(l.v1.X, l.v2.X)
			for i := x1; i <= x2; i++ {
				s[l.v2.Y][i-rec.Ori.X] = rock
			}
		}
	}
	return &caveSlice{
		s:      s,
		height: height,
		width:  rec.Width,
		src:    shared.Coord{X: 500 - rec.Ori.X},
	}
}

func parseLines(path string) (lines []line, rec shared.Rectangle) {
	fileLines := utils.LoadFileLines(path)
	l := len(fileLines) - 1
	list := collections.NewArrayListOf[line]()
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for i := 0; i < l; i++ {
		fileLine := fileLines[i]
		lineStrs := strings.Split(fileLine, " -> ")
		for j := 1; j < len(lineStrs); j++ {
			v1 := shared.ParseCoord(lineStrs[j-1])
			v2 := shared.ParseCoord(lineStrs[j])
			var isVertical bool
			if v1.X == v2.X {
				isVertical = true
			} else if v1.Y == v2.Y {
				isVertical = false
			} else {
				panic("can only have vertical or horizontal lines")
			}
			list.Add(line{v1, v2, isVertical})
			minX = utils.MinOf(minX, v1.X, v2.X)
			maxX = utils.MaxOf(maxX, v1.X, v2.X)
			minY = utils.MinOf(minY, v1.Y, v2.Y)
			maxY = utils.MaxOf(maxY, v1.Y, v2.Y)
		}
	}

	return utils2.SliceFrom[line](list), shared.Rectangle{
		Ori:    shared.Coord{X: minX, Y: minY},
		Width:  maxX - minX + 1,
		Height: maxY - minY + 1,
	}
}

func (cs *caveSlice) String() string {
	var b strings.Builder
	for i := 0; i < cs.height; i++ {
		for j := 0; j < cs.width; j++ {
			c := cs.s[i][j]
			if c == rock {
				b.WriteByte('#')
			} else if c == air {
				b.WriteByte('.')
			} else if c == sand {
				b.WriteByte('o')
			} else {
				panic("unknown char")
			}
		}
		b.WriteByte('\n')
	}
	return b.String()
}
