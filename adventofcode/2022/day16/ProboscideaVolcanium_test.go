package day16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePipeGraph(t *testing.T) {
	pg := parsePipeGraph("./test-input")
	assert.Equal(t, "AA", pg.start.name)
	assert.Equal(t, 0, pg.start.rate)

	assert.Equal(t, 10, len(pg.valves))

	assert.Equal(t, []string{"DD", "II", "BB"}, pg.start.next)

	ff := pg.valves["FF"]
	assert.Equal(t, 0, ff.rate)
	assert.Equal(t, []string{"EE", "GG"}, ff.next)
}
