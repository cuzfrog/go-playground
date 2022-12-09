package day9

import (
	"github.com/cuzfrog/go-playground/adventofcode/2022/shared"
	"github.com/cuzfrog/go-playground/utils"
	"github.com/cuzfrog/tgods/collections"
	"github.com/cuzfrog/tgods/types"
	cutils "github.com/cuzfrog/tgods/utils"
)

/*
The coordination system
Y
3
2
1
0 1 2 3 X
*/

type Direction byte

const (
	Up    Direction = 'U'
	Right Direction = 'R'
	Down  Direction = 'D'
	Left  Direction = 'L'
)

type step struct {
	di  Direction
	cnt int
}

type knot struct {
	cur   shared.Coord
	route types.List[shared.Coord]
}

func performSteps(steps []step) (*knot, *knot) {
	head, tail := newKnot(), newKnot()
	for _, s := range steps {
		for i := 0; i < s.cnt; i++ {
			head.move(s.di)
			tail.follow(head)
		}
	}
	return head, tail
}

func (k *knot) move(di Direction) {
	if di == Up {
		k.cur.Y++
	} else if di == Right {
		k.cur.X++
	} else if di == Down {
		k.cur.Y--
	} else if di == Left {
		k.cur.X--
	} else {
		panic("Unknown direction")
	}
	k.route.Add(k.cur)
}

func (k *knot) follow(h *knot) {
	xDif := h.cur.X - k.cur.X
	yDif := h.cur.Y - k.cur.Y
	distance2 := xDif*xDif + yDif*yDif
	if distance2 <= 2 {
		return
	}

	if distance2 == 4 {
		k.cur.X += xDif / 2
		k.cur.Y += yDif / 2
	} else if distance2 == 5 {
		k.cur.X += xDif / utils.Abs(xDif)
		k.cur.Y += yDif / utils.Abs(yDif)
	} else {
		panic("impossible distance power2")
	}
	k.route.Add(k.cur)
}

func (k *knot) countNoDuplicatePosition() int {
	set := collections.NewHashSetC[shared.Coord]()
	cutils.AddAll[shared.Coord](k.route, set)
	return set.Size()
}

func newKnot() *knot {
	s := shared.Coord{}
	k := knot{s, collections.NewArrayListOf[shared.Coord]()}
	k.route.Add(s)
	return &k
}

func parseSteps(path string) []step {
	lines := utils.LoadFileLines(path)
	l := len(lines) - 1
	steps := make([]step, l)
	for i := 0; i < l; i++ {
		diStr, cntStr := utils.SplitString2(lines[i], " ")
		steps[i] = step{Direction(diStr[0]), utils.StrToInt(cntStr)}
	}
	return steps
}
