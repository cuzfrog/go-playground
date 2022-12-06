package day4

import (
	"fmt"
	"github.com/cuzfrog/go-playground/utils"
	"strings"
)

type assignment struct {
	start int
	end   int
}

type pair struct {
	elf1 assignment
	elf2 assignment
}

func (p *pair) String() string {
	return fmt.Sprintf("%d-%d,%d-%d", p.elf1.start, p.elf1.end, p.elf2.start, p.elf2.end)
}

func (p *pair) coverTheOther() bool {
	return (p.elf1.start <= p.elf2.start && p.elf1.end >= p.elf2.end) || (p.elf1.start >= p.elf2.start && p.elf1.end <= p.elf2.end)
}

func (p *pair) overlap() bool {
	res := p.coverTheOther() || (p.elf1.start >= p.elf2.start && p.elf1.start <= p.elf2.end) || (p.elf1.end >= p.elf2.start && p.elf1.end <= p.elf2.end)
	//println(fmt.Sprintf("%s %v", p, res))
	return res
}

func parsePairs(path string) []pair {
	lines := utils.LoadFileLines(path)
	l := len(lines) - 1
	pairs := make([]pair, l)
	for i := 0; i < l; i++ {
		pairs[i] = parsePair(lines[i])
	}
	return pairs
}

func parsePair(line string) pair {
	ss := strings.Split(line, ",")
	return pair{parseCover(ss[0]), parseCover(ss[1])}
}

func parseCover(r string) assignment {
	cc := strings.Split(r, "-")
	as := assignment{utils.StrToInt(cc[0]), utils.StrToInt(cc[1])}
	if as.start > as.end {
		panic("invalid start and end")
	}
	return as
}
