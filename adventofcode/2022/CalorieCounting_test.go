package _2022

import "testing"

func TestFindMaxCalorie(t *testing.T) {
	elves := parseElves("./day1input")
	elves.Sort(func(e1, e2 *Elf) bool { return e1.totalCalories > e2.totalCalories })
	elf1, _ := elves.Head()
	println(elf1.totalCalories)
}

func TestFindMax3Calories(t *testing.T) {
	elves := parseElves("./day1input")
	elves.Sort(func(e1, e2 *Elf) bool { return e1.totalCalories > e2.totalCalories })
	sum := 0
	for i := 0; i < 3; i++ {
		elf, _ := elves.Get(i)
		sum += elf.totalCalories
	}
	println(sum)
}
