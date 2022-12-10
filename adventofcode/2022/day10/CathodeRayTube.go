package day10

import "github.com/cuzfrog/go-playground/utils"

type instruction struct {
	cost int
	v    int
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
