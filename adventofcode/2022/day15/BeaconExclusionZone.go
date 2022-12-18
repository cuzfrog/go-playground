package day15

import (
	"fmt"
	"github.com/cuzfrog/go-playground/adventofcode/2022/shared"
	"github.com/cuzfrog/go-playground/utils"
	"math"
	"regexp"
)

const (
	unknown byte = iota
	sensor
	beacon
	noBeacon
)

func sprintFn(v byte) string {
	var c byte
	if v == unknown {
		c = '.'
	} else if v == sensor {
		c = 'S'
	} else if v == beacon {
		c = 'B'
	} else if v == noBeacon {
		c = '#'
	} else {
		panic("invalid value")
	}
	return fmt.Sprintf("%c", c)
}

type pair struct {
	se shared.Coord
	be shared.Coord
}

func scanUpdateChart(ch *shared.Chart[byte], pairs []pair) {
	for _, p := range pairs {
		p.markNoBeacon(ch)
	}
}

func (p pair) markNoBeacon(ch *shared.Chart[byte]) {
	dis := manhattanDistance(p.se, p.be)
	ch.Each(func(x, y int, v byte) {
		if manhattanDistance(shared.Coord{X: x, Y: y}, p.se) <= dis && v == unknown {
			ch.Put(x, y, noBeacon)
		}
	})
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

func manhattanDistance(c1, c2 shared.Coord) int {
	return utils.Abs(c1.X-c2.X) + utils.Abs(c1.Y-c2.Y)
}
