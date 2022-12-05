package day3

import (
	"fmt"
	"github.com/cuzfrog/go-playground/utils"
	"github.com/cuzfrog/tgods/collections"
	"github.com/cuzfrog/tgods/transform"
	"github.com/cuzfrog/tgods/types"
	utils2 "github.com/cuzfrog/tgods/utils"
)

func sumPriority(items types.List[uint8]) int {
	scores := collections.NewArrayListOfSize[int](items.Size())
	transform.MapTo[uint8, int](items, scores, toPriority)
	sum := transform.Reduce[int, int](scores, 0, func(acc int, next int) int { return acc + next })
	return sum
}

func findAllShared(path string) types.List[uint8] {
	lines := utils.LoadFileLines(path)
	l := len(lines) - 1
	all := collections.NewArrayListOfEq[types.Set[uint8]](l, nil)
	for i := 0; i < l; i++ {
		sack := lines[i]
		items := findShared(sack)
		all.Add(items)
		itemsStr := transform.Reduce[uint8, string](items, "", func(acc string, next uint8) string { return fmt.Sprintf("%s %c", acc, next) })
		println(itemsStr)
	}
	flatAll := collections.NewArrayListOf[uint8]()
	transform.FlatMapTo[types.Set[uint8], uint8](all, flatAll, func(s types.Set[uint8]) []uint8 { return utils2.SliceFrom[uint8](s) })
	return flatAll
}

func findShared(sack string) types.Set[uint8] {
	l := len(sack)
	if l&1 > 0 {
		panic("sack contains odd number of items")
	}

	chars := collections.NewHashSetOfNum[uint8]()
	for i := 0; i < l>>1; i++ {
		for j := l >> 1; j < l; j++ {
			if sack[i] == sack[j] {
				chars.Add(sack[i]) // TODO deduplication
			}
		}
	}
	return chars
}

func toPriority(c uint8) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	} else if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 27)
	}
	panic(fmt.Sprintf("invalid char: %c", c))
}
