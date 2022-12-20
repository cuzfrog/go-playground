package shared

import "github.com/cuzfrog/go-playground/utils"

func LoadInput(path string) (lines []string, length int) {
	rows := utils.LoadFileLines(path)
	length = len(rows) - 1
	return rows[:length], length
}
