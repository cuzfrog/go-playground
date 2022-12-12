package day12

import (
	"fmt"
	"github.com/cuzfrog/go-playground/adventofcode/2022/shared"
	"github.com/cuzfrog/go-playground/utils"
	"github.com/cuzfrog/tgods/collections"
	"github.com/cuzfrog/tgods/transform"
	"github.com/cuzfrog/tgods/types"
	utils2 "github.com/cuzfrog/tgods/utils"
	"strings"
)

type step struct {
	value int8
	coord shared.Coord
	prev  *step
}

func (s *step) String() string {
	return fmt.Sprintf("%v", s.coord)
}

func (s *step) Route() string {
	cur := s
	stack := collections.NewLinkedListStack[*step]()
	for cur != nil {
		stack.Push(cur)
		cur = cur.prev
	}
	b := new(strings.Builder)
	it := stack.Iterator()
	for it.Next() {
		b.WriteString(fmt.Sprintf("%c(%v)", it.Value().value, it.Value().String()))
		b.WriteByte('>')
	}
	return b.String()
}

type theMap struct {
	heights [][]int8
	start   shared.Coord
	end     shared.Coord
	rowCnt  int
	colCnt  int
}

func moveCnt(m *theMap) (int, *step) {
	visited := collections.NewHashSetC[shared.Coord]()
	visited.Add(m.start)
	curSteps := collections.NewArrayListOf(&step{m.height(m.start), m.start, nil})
	cnt := 0
	for !visited.Contains(m.end) {
		nextSteps := collections.NewArrayListOf[*step]()
		transform.FlatMapTo[*step, *step](curSteps, nextSteps, func(cur *step) []*step { return nextMoves(m, cur, visited) })
		if nextSteps.Size() == 0 {
			panic(fmt.Sprintf("no way to go, curs: %v", *utils2.StringFrom[*step](curSteps)))
		}
		curSteps = nextSteps
		cnt++
	}
	it := curSteps.Iterator()
	var endStep *step
	for it.Next() {
		if it.Value().coord.Equal(m.end) {
			endStep = it.Value()
		}
	}
	return cnt, endStep
}

func nextMoves(m *theMap, cur *step, visited types.Set[shared.Coord]) []*step {
	mvs := collections.NewArrayListOfSize[*step](4)
	n, ok := moveUp(m, cur.coord)
	if ok && !visited.Contains(n) {
		mvs.Add(&step{m.height(n), n, cur})
	}
	n, ok = moveRight(m, cur.coord)
	if ok && !visited.Contains(n) {
		mvs.Add(&step{m.height(n), n, cur})
	}
	n, ok = moveDown(m, cur.coord)
	if ok && !visited.Contains(n) {
		mvs.Add(&step{m.height(n), n, cur})
	}
	n, ok = moveLeft(m, cur.coord)
	if ok && !visited.Contains(n) {
		mvs.Add(&step{m.height(n), n, cur})
	}
	transform.MapTo[*step, shared.Coord](mvs, visited, func(s *step) shared.Coord { return s.coord })
	return utils2.SliceFrom[*step](mvs)
}

func loadMap(path string) *theMap {
	lines := utils.LoadFileLines(path)
	l := len(lines) - 1
	colCnt := len(lines[0])
	m := theMap{make([][]int8, l), shared.Coord{}, shared.Coord{}, l, colCnt}
	maxc := int8('a')
	for i := 0; i < l; i++ {
		line := lines[i]
		m.heights[i] = make([]int8, colCnt)
		for j := 0; j < colCnt; j++ {
			c := int8(line[j])
			if c > maxc {
				maxc = c
			}
			m.heights[i][j] = c
			if line[j] == 'S' {
				m.start.X = j
				m.start.Y = i
			} else if line[j] == 'E' {
				m.end.X = j
				m.end.Y = i
			}
		}
	}
	m.heights[m.start.Y][m.start.X] = 'a'
	m.heights[m.end.Y][m.end.X] = maxc
	return &m
}

func (m *theMap) height(coord shared.Coord) int8 {
	return m.heights[coord.Y][coord.X]
}
func moveLeft(m *theMap, cur shared.Coord) (shared.Coord, bool) {
	left := shared.Coord{X: cur.X - 1, Y: cur.Y}
	if cur.X == 0 {
		return left, false
	}
	return left, m.height(left)-m.height(cur) <= 1
}
func moveRight(m *theMap, cur shared.Coord) (shared.Coord, bool) {
	right := shared.Coord{X: cur.X + 1, Y: cur.Y}
	if cur.X == m.colCnt-1 {
		return right, false
	}
	return right, m.height(right)-m.height(cur) <= 1
}
func moveUp(m *theMap, cur shared.Coord) (shared.Coord, bool) {
	up := shared.Coord{X: cur.X, Y: cur.Y - 1}
	if cur.Y == 0 {
		return up, false
	}
	return up, m.height(up)-m.height(cur) <= 1
}
func moveDown(m *theMap, cur shared.Coord) (shared.Coord, bool) {
	down := shared.Coord{X: cur.X, Y: cur.Y + 1}
	if cur.Y == m.rowCnt-1 {
		return down, false
	}
	return down, m.height(down)-m.height(cur) <= 1
}
