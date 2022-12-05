package day2

import (
	"fmt"
	"github.com/cuzfrog/go-playground/utils"
	"strings"
)

type strategy int8

const (
	rock     strategy = 1
	paper    strategy = 2
	scissors strategy = 3
)

type round struct {
	opponent strategy
	me       strategy
}

func score(r round) int {
	s := int(r.me)
	if r.opponent == r.me {
		s += 3
	} else if (r.opponent == rock && r.me == paper) || (r.opponent == paper && r.me == scissors) || (r.opponent == scissors && r.me == rock) {
		s += 6
	}
	println(fmt.Sprintf("Opponent: %s ,Me: %s, score: %d", r.opponent, r.me, s))
	return s
}

func (s strategy) String() string {
	if s == rock {
		return "rock"
	} else if s == paper {
		return "paper"
	} else if s == scissors {
		return "scissors"
	}
	panic("invalid strategy")
}

func parseRounds(inputPath string) []round {
	lines := utils.LoadFileLines(inputPath)
	rounds := make([]round, len(lines)-1)
	for i, line := range lines {
		if line != "" {
			cols := strings.Split(line, " ")
			rounds[i] = round{parseStrategy(cols[0]), parseStrategy(cols[1])}
		}
	}
	return rounds
}

func parseStrategy(input string) strategy {
	if input == "A" || input == "X" {
		return rock
	}
	if input == "B" || input == "Y" {
		return paper
	}
	if input == "C" || input == "Z" {
		return scissors
	}
	panic(fmt.Sprintf(`invalid input: '%s'`, input))
}
