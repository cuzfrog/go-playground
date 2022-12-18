package day15

import (
	"github.com/cuzfrog/go-playground/adventofcode/2022/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parsePairs(t *testing.T) {
	rec, pairs := parseChart("./test-input")
	assert.Equal(t, 14, len(pairs))
	assert.Equal(t, shared.NewCoord(2, 18), pairs[0].se)
	assert.Equal(t, shared.NewCoord(-2, 15), pairs[0].be)
	assert.Equal(t, shared.NewCoord(14, 3), pairs[12].se)
	assert.Equal(t, shared.NewCoord(15, 3), pairs[12].be)

	assert.Equal(t, shared.NewCoord(-8, -10), rec.Ori)
}

func TestCheckBeacon(t *testing.T) {
	_, pairs := parseChart("./test-input")
	assert.Equal(t, unknown, testBeacon(-2, 9, pairs))
	assert.Equal(t, noBeacon, testBeacon(0, 7, pairs))
	assert.Equal(t, noBeacon, testBeacon(-1, 7, pairs))
	assert.Equal(t, unknown, testBeacon(-2, 7, pairs))
	assert.Equal(t, sensor, testBeacon(8, 7, pairs))
	assert.Equal(t, unknown, testBeacon(-3, 10, pairs))
	assert.Equal(t, noBeacon, testBeacon(-2, 10, pairs))
	assert.Equal(t, beacon, testBeacon(2, 10, pairs))
	assert.Equal(t, unknown, testBeacon(14, 11, pairs))
}

func TestCountNoBeacon(t *testing.T) {
	rec, pairs := parseChart("./test-input")
	cnt := countNoBeaconOnRow(10, rec, pairs)
	assert.Equal(t, 26, cnt)
}

func TestSearchUnknown(t *testing.T) {
	_, pairs := parseChart("./test-input")
	//pos := searchUnknown(0, 11, 14, 11, pairs)

	pos := searchUnknown(0, 0, 20, 20, pairs)
	assert.Equal(t, shared.NewCoord(14, 11), pos)
	freq := pos.X*4000000 + pos.Y
	assert.Equal(t, 56000011, freq)
}

func TestSolution1(t *testing.T) {
	rec, pairs := parseChart("./input")
	cnt := countNoBeaconOnRow(2000000, rec, pairs)
	println(cnt)
	assert.Equal(t, 5176944, cnt)
}

func TestSolution2(t *testing.T) {
	rec, pairs := parseChart("./input")
	println(rec.String())

	pos := searchUnknown(0, 0, 4000000, 4000000, pairs)
	freq := pos.X*4000000 + pos.Y
	println(freq)
	assert.Equal(t, 13350458933732, freq)
}
