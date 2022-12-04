package _2022

import (
	"github.com/cuzfrog/tgods/collections"
	"github.com/cuzfrog/tgods/types"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	num           int
	snacks        types.List[int]
	totalCalories int
}

func elfEqual(elf1 *Elf, elf2 *Elf) bool {
	return elf1.num == elf2.num
}

func parseElves(inputPath string) types.ArrayList[*Elf] {
	content := loadInput(inputPath)
	lines := strings.Split(content, "\n")
	elves := collections.NewArrayListOfEq[*Elf](100, elfEqual)

	num := 1
	cur := &Elf{1, collections.NewArrayListOf[int](), 0}
	for _, line := range lines {
		if line == "" {
			elves.Add(cur)
			num++
			cur = &Elf{num, collections.NewArrayListOf[int](), 0}
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			cur.snacks.Add(num)
			cur.totalCalories += num
		}
	}
	return elves
}

func loadInput(path string) string {
	dat, err := os.ReadFile("./day1input")
	if err != nil {
		panic(err)
	}
	return string(dat)
}
