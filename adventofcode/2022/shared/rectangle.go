package shared

import "fmt"

type Rectangle struct {
	Ori    Coord
	Width  int
	Height int
}

func (rec *Rectangle) String() string {
	return fmt.Sprintf("Rec[%v, %dx%d]", rec.Ori, rec.Width, rec.Height)
}
