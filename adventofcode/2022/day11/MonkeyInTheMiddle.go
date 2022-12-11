package day11

import (
	"github.com/cuzfrog/go-playground/utils"
	"github.com/cuzfrog/tgods/collections"
	"github.com/cuzfrog/tgods/types"
	utils2 "github.com/cuzfrog/tgods/utils"
	"strings"
)

type monkey struct {
	items        types.Queue[int]
	ops          func(old int) int
	testDividant int
	tgtT         int //monkey index
	tgtF         int //monkey index
}

func parseMonkeys(path string) []*monkey {
	lines := utils.LoadFileLines(path)
	l := len(lines) - 1
	monkeys := collections.NewArrayListOfSize[*monkey](10)
	var m *monkey
	for i := 0; i < l; i++ {
		line := lines[i]
		if strings.HasPrefix(line, "Monkey") {
			_, idStr := utils.SplitString2(line[:len(line)-1], " ")
			id := utils.StrToInt(idStr)
			m = &monkey{items: collections.NewArrayListQueue[int]()}
			monkeys.Set(id, m)
		} else if strings.HasPrefix(line, "  Starting items:") {
			itemsStr := strings.Split(line[18:], ", ")
			for _, itemStr := range itemsStr {
				m.items.Enqueue(utils.StrToInt(itemStr))
			}
		} else if strings.HasPrefix(line, "  Operation:") {
			_, opsStr := utils.SplitString2(line, "old ")
			operant := opsStr[2:]
			var opFn func(a, b int) int
			if opsStr[0] == '+' {
				opFn = func(a, b int) int { return a + b }
			} else if opsStr[0] == '*' {
				opFn = func(a, b int) int { return a * b }
			} else {
				panic("unknown operator")
			}
			if operant == "old" {
				m.ops = func(old int) int { return opFn(old, old) }
			} else {
				num := utils.StrToInt(operant)
				m.ops = func(old int) int { return opFn(old, num) }
			}
		} else if strings.HasPrefix(line, "  Test: divisible by ") {
			m.testDividant = utils.StrToInt(line[21:])
		} else if strings.HasPrefix(line, "    If true: throw to monkey ") {
			m.tgtT = utils.StrToInt(line[29:])
		} else if strings.HasPrefix(line, "    If false: throw to monkey ") {
			m.tgtF = utils.StrToInt(line[30:])
		} else if line == "" {
			// skip
		} else {
			panic("unknown input line")
		}
	}
	return utils2.SliceFrom[*monkey](monkeys)
}
