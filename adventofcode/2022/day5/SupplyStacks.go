package day5

import (
	"fmt"
	"github.com/cuzfrog/go-playground/utils"
	"github.com/cuzfrog/tgods/collections"
	"github.com/cuzfrog/tgods/types"
	"regexp"
)

type stacks []types.Stack[uint8]

type move struct {
	cnt int
	src int
	tgt int
}

func (ss stacks) peerTop() string {
	chars := make([]uint8, len(ss))
	for i, s := range ss {
		item, ok := s.Peek()
		if !ok {
			panic("no item left in stack")
		}
		chars[i] = item
	}
	return string(chars)
}

func (ss stacks) performMove(mv move) {
	for i := 0; i < mv.cnt; i++ {
		item, ok := ss[mv.src-1].Pop()
		if !ok {
			panic("no item left in stack")
		}
		ss[mv.tgt-1].Push(item)
	}
}

var moveRegex = regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")

func parseMoves(path string) []move {
	lines := utils.LoadFileLines(path)
	l := len(lines) - 1
	moves := make([]move, l)
	for i := 0; i < l; i++ {
		match := moveRegex.FindStringSubmatch(lines[i])
		moves[i] = move{utils.StrToInt(match[1]), utils.StrToInt(match[2]), utils.StrToInt(match[3])}
	}
	return moves
}

func parseStacks(path string) stacks {
	lines := utils.LoadFileLines(path)
	l := len(lines) - 1
	numLine := lines[l-1]
	scnt := int(numLine[len(numLine)-1] - '0')
	println(fmt.Sprintf("stack count: %d", scnt))

	ss := make(stacks, scnt)
	for i := 0; i < scnt; i++ {
		ss[i] = collections.NewLinkedListStack[uint8]()
	}

	for i := l - 2; i >= 0; i-- {
		row := lines[i]
		for j := 0; j < scnt; j++ {
			colIndex := j*4 + 1
			if len(row) > colIndex {
				item := row[colIndex]
				if utils.IsUpperCaseLetter(item) {
					ss[j].Push(item)
				} else if item != ' ' {
					panic(fmt.Sprintf("parsed wrong char: %c", item))
				}
			}
		}
	}

	return ss
}
