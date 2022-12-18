package day15

import (
	"github.com/cuzfrog/go-playground/adventofcode/2022/shared"
	"github.com/cuzfrog/go-playground/utils"
	"math"
	"regexp"
)

const (
	empty byte = iota
	sensor
	beacon
)

type pair struct {
	se shared.Coord
	be shared.Coord
}

func parseChart(path string) (*shared.Chart[byte], []pair) {
	regex := regexp.MustCompile("Sensor at x=(-?\\d+), y=(-?\\d+): closest beacon is at x=(-?\\d+), y=(-?\\d+)")
	lines := utils.LoadFileLines(path)
	l := len(lines) - 1
	pairs := make([]pair, l)

	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for i := 0; i < l; i++ {
		match := regex.FindStringSubmatch(lines[i])
		se := shared.Coord{X: utils.StrToInt(match[1]), Y: utils.StrToInt(match[2])}
		be := shared.Coord{X: utils.StrToInt(match[3]), Y: utils.StrToInt(match[4])}
		pairs[i] = pair{se, be}

		minX = utils.MinOf(minX, se.X, be.X)
		maxX = utils.MaxOf(maxX, se.X, be.X)
		minY = utils.MinOf(minY, se.Y, be.Y)
		maxY = utils.MaxOf(maxY, se.Y, be.Y)
	}

	width := maxX - minX + 1
	height := maxY - minY + 1
	ch := shared.NewChart[byte](minX, minY, width, height)

	for i := 0; i < l; i++ {
		p := pairs[i]
		ch.Put(p.se.X, p.se.Y, sensor)
		ch.Put(p.be.X, p.be.Y, beacon)
	}

	return ch, pairs
}
