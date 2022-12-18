package shared

import (
	"fmt"
	"github.com/cuzfrog/go-playground/utils"
)

type Coord struct {
	X int
	Y int
}

func (c Coord) Hash() uint {
	return uint(c.X*17 + c.Y)
}
func (c Coord) Equal(o Coord) bool {
	return c.X == o.X && c.Y == o.Y
}
func (c Coord) String() string {
	return fmt.Sprintf("%d.%d", c.X, c.Y)
}
func ParseCoord(s string) Coord {
	v1, v2 := utils.SplitString2(s, ",")
	return Coord{utils.StrToInt(v1), utils.StrToInt(v2)}
}
func NewCoord(x, y int) Coord {
	return Coord{x, y}
}
