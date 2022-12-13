package day13

import (
	"github.com/cuzfrog/go-playground/utils"
)

type signal struct {
	isList bool
	v      int
	values []*signal
}

type pair struct {
	left  *signal
	right *signal
}

type span struct {
	start int // inclusive
	end   int // exclusive
}

func (s *span) len() int {
	return s.end - s.start
}

func findOrder(p pair) int {
	if p.left.isList && p.right.isList {
		ll := len(p.left.values)
		rl := len(p.right.values)
		l := utils.Min(ll, rl)
		for i := 0; i < l; i++ {
			ls := p.left.values[i]
			rs := p.right.values[i]
			res := findOrder(pair{ls, rs})
			if res < 0 {
				return -1
			} else if res > 0 {
				return 1
			}
		}
		if ll < rl {
			return -1
		} else if ll > rl {
			return 1
		} else {
			return 0
		}
	} else if !p.left.isList && !p.right.isList {
		if p.left.v < p.right.v {
			return -1
		} else if p.left.v > p.right.v {
			return 1
		} else {
			return 0
		}
	} else if !p.left.isList && p.right.isList {
		return findOrder(pair{&signal{true, -1, []*signal{p.left}}, p.right})
	} else if p.left.isList && !p.right.isList {
		return findOrder(pair{p.left, &signal{true, -1, []*signal{p.right}}})
	} else {
		panic("")
	}
}

func parseSignals(path string) []pair {
	lines := utils.LoadFileLines(path)
	l := len(lines) - 1
	pairs := make([]pair, (l+1)/3)
	pIdx := 0
	for i := 0; i < l; i += 3 {
		p := pair{}
		left := lines[i]
		p.left = parseSignal([]byte(left), span{0, len(left)})
		right := lines[i+1]
		p.right = parseSignal([]byte(right), span{0, len(right)})
		pairs[pIdx] = p
		pIdx++
	}
	return pairs
}

func parseSignal(line []byte, sp span) *signal {
	if sp.len() <= 0 {
		panic("empty input")
	}
	if line[sp.start] == '[' {
		spans := splitElems(line, span{sp.start + 1, sp.end - 1})
		sig := &signal{true, -1, make([]*signal, len(spans))}
		for i, s := range spans {
			sig.values[i] = parseSignal(line, s)
		}
		return sig
	} else {
		num := utils.StrToInt(string(line[sp.start:sp.end])) // TODO avoid string conversion
		return &signal{false, num, nil}
	}
}

var buf = make([]span, 64)

func splitElems(line []byte, sp span) []span {
	if sp.len() <= 0 {
		return []span{}
	}
	level := 0
	curIdx := 0
	cur := span{sp.start, sp.start}
	for i := sp.start; i < sp.end+1; i++ {
		var c byte
		if i < sp.end {
			c = line[i]
		} else {
			c = ','
		}
		if c == '[' {
			level++
		} else if c == ']' {
			level--
		} else if c == ',' && level == 0 {
			buf[curIdx] = cur
			curIdx++
			cur = span{i + 1, i + 1}
			continue
		}
		cur.end++
	}
	res := make([]span, curIdx)
	copy(res, buf)
	return res
}

func signalLess(a, b *signal) bool {
	return findOrder(pair{a, b}) < 0
}

func pairToSignals(p pair) []*signal {
	return []*signal{p.left, p.right}
}
