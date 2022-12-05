package utils

import (
	"fmt"
	"os"
	"strings"
)

func LoadFileContent(path string) string {
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func LoadFileLines(path string) []string {
	lines := strings.Split(LoadFileContent(path), "\n")
	if len(lines) > 0 && strings.Contains(lines[0], "\r") {
		panic(fmt.Sprintf("only support LF line ending, path: %s", path))
	}
	return lines
}
