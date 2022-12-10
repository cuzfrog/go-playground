package day10

import (
	"bytes"
	"github.com/cuzfrog/go-playground/utils"
)

type instruction struct {
	cost int
	v    int
}

type sprite struct {
	pos   int
	cycle int
}

type frame [][]byte

func (f *frame) Sprint() string {
	b := bytes.NewBuffer(make([]byte, 0, 240))
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			p := (*f)[i][j]
			b.WriteByte(p)
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func parseSprites(instructions []instruction) []sprite {
	cycle := 1
	pos := 1
	sprites := make([]sprite, 240)
	for _, ins := range instructions {
		for i := 0; i < ins.cost; i++ {
			sprites[cycle-1] = sprite{pos, cycle}
			cycle += 1
			if i == ins.cost-1 {
				pos += ins.v
			}
		}
	}
	return sprites
}

func scan(sprites []sprite) frame {
	cycle := 1
	f := make([][]byte, 6)
	for i := 0; i < 6; i++ {
		f[i] = make([]byte, 40)
		line := f[i]
		for j := 0; j < 40; j++ {
			s := sprites[cycle-1]
			if utils.Abs(s.pos-j) <= 1 {
				line[j] = '#'
			} else {
				line[j] = '.'
			}
			cycle++
		}
	}
	return f
}

func sumSignalStrength(instructions []instruction) int {
	cycle := 1
	reg := 1
	strength := 0
	for _, ins := range instructions {
		for i := 0; i < ins.cost; i++ {
			cycle += 1
			if i == ins.cost-1 {
				reg += ins.v
			}
			if cycle == 20 || (cycle-20)%40 == 0 {
				strength += cycle * reg
			}
		}
	}
	return strength
}

func parseInstructions(path string) []instruction {
	lines := utils.LoadFileLines(path)
	l := len(lines) - 1
	instructions := make([]instruction, l)
	for i := 0; i < l; i++ {
		line := lines[i]
		if line == "noop" {
			instructions[i] = instruction{1, 0}
		} else {
			_, vStr := utils.SplitString2(line, " ")
			instructions[i] = instruction{2, utils.StrToInt(vStr)}
		}
	}
	return instructions
}
