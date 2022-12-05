package day1

import (
	"github.com/cuzfrog/go-playground/utils"
	"github.com/cuzfrog/tgods/collections"
	"github.com/cuzfrog/tgods/types"
	"strconv"
)

type elf struct {
	num           int
	snacks        types.List[int]
	totalCalories int
}

func elfEqual(elf1 *elf, elf2 *elf) bool {
	return elf1.num == elf2.num
}

func parseElves(inputPath string) types.ArrayList[*elf] {
	lines := utils.LoadFileLines(inputPath)
	elves := collections.NewArrayListOfEq[*elf](100, elfEqual)

	num := 1
	cur := &elf{1, collections.NewArrayListOf[int](), 0}
	for _, line := range lines {
		if line == "" {
			elves.Add(cur)
			num++
			cur = &elf{num, collections.NewArrayListOf[int](), 0}
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
