package day16

import (
	"github.com/cuzfrog/go-playground/adventofcode/2022/shared"
	"github.com/cuzfrog/go-playground/utils"
	"regexp"
	"strings"
)

type valve struct {
	name string
	rate int
	next []string
}

type pipeGraph struct {
	valves map[string]*valve
	start  *valve
}

func parsePipeGraph(path string) *pipeGraph {
	regex := regexp.MustCompile("Valve (\\w+) has flow rate=(\\d+); tunnels? leads? to valves? ([,\\s\\w]+)")
	lines, l := shared.LoadInput(path)
	valves := make(map[string]*valve, l)
	var start *valve
	for i := 0; i < l; i++ {
		match := regex.FindStringSubmatch(lines[i])
		nextNames := strings.Split(match[3], ", ")
		v := &valve{match[1], utils.StrToInt(match[2]), make([]string, len(nextNames))}
		for j := 0; j < len(nextNames); j++ {
			v.next[j] = nextNames[j]
		}
		if i == 0 {
			start = v
		}
		valves[v.name] = v
	}
	return &pipeGraph{valves, start}
}
