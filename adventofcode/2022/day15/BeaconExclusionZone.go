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
	se  shared.Coord
	be  shared.Coord
	dis int
}

func countNoBeaconOnRow(y int, rec shared.Rectangle, pairs []pair) int {
	x0 := rec.Ori.X
	x1 := x0 + rec.Width
	cnt := 0
	for i := x0; i < x1; i++ {
		res := testBeacon(i, y, pairs)
		if res == noBeacon {
			cnt++
		}
	}
	return cnt
}

func searchUnknown(x0, y0, x1, y1 int, pairs []pair) shared.Coord {
	y := x0
	for y <= y1 {
		x := x0
		for x <= x1 {
			b, d, se := testBeaconAndDis(x, y, pairs)
			if b == unknown { // d < 0
				return shared.NewCoord(x, y)
			}
			if x < se.X {
				x = x + (se.X-x)*2 + 1
			} else if x >= se.X {
				x = x + d + 1
			}
		}
		//println(y)
		y++
	}
	panic("no possible")
}

func testBeacon(x, y int, pairs []pair) byte {
	b, _, _ := testBeaconAndDis(x, y, pairs)
	return b
}

// return v, disDiff, sensorPos
func testBeaconAndDis(x, y int, pairs []pair) (byte, int, shared.Coord) {
	co := shared.Coord{X: x, Y: y}
	res := unknown
	d := math.MaxInt
	for _, p := range pairs {
		d = p.dis - manhattanDistance(co, p.se)
		if d >= 0 {
			if co == p.be {
				return beacon, d, p.se
			} else if co == p.se {
				return sensor, d, p.se
			} else {
				return noBeacon, d, p.se
			}
		} else {
			continue
		}
	}
	return res, d, shared.Coord{}
}

func parseChart(path string) (shared.Rectangle, []pair) {
	regex := regexp.MustCompile("Sensor at x=(-?\\d+), y=(-?\\d+): closest beacon is at x=(-?\\d+), y=(-?\\d+)")
	lines := utils.LoadFileLines(path)
	l := len(lines) - 1
	pairs := make([]pair, l)

	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for i := 0; i < l; i++ {
		match := regex.FindStringSubmatch(lines[i])
		se := shared.Coord{X: utils.StrToInt(match[1]), Y: utils.StrToInt(match[2])}
		be := shared.Coord{X: utils.StrToInt(match[3]), Y: utils.StrToInt(match[4])}
		dis := manhattanDistance(se, be)
		pairs[i] = pair{se, be, dis}

		minX = utils.MinOf(minX, se.X-dis, be.X)
		maxX = utils.MaxOf(maxX, se.X+dis, be.X)
		minY = utils.MinOf(minY, se.Y-dis, be.Y)
		maxY = utils.MaxOf(maxY, se.Y+dis, be.Y)
	}

	width := maxX - minX + 1
	height := maxY - minY + 1
	return shared.Rectangle{Ori: shared.Coord{X: minX, Y: minY}, Width: width, Height: height}, pairs
}

func manhattanDistance(c1, c2 shared.Coord) int {
	return utils.Abs(c1.X-c2.X) + utils.Abs(c1.Y-c2.Y)
}
