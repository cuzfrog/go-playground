package shared

import "fmt"

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
