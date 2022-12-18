package shared

import "strings"

type Chart[T comparable] struct {
	data   []Row[T]
	ori    Coord
	width  int
	height int
}

type Row[T comparable] []T

func NewChart[T comparable](oriX, oriY int, w, h int) *Chart[T] {
	data := make([]Row[T], h)
	for i := 0; i < h; i++ {
		data[i] = make([]T, w)
	}
	return &Chart[T]{data, Coord{X: oriX, Y: oriY}, w, h}
}

func (ch *Chart[T]) Size() (width, height int) {
	return ch.width, ch.height
}

func (ch *Chart[T]) Origin() Coord {
	return ch.ori
}

func (ch *Chart[T]) Get(x, y int) T {
	return ch.data[y-ch.ori.Y][x-ch.ori.X]
}

func (ch *Chart[T]) GetRow(y int) Row[T] {
	return ch.data[y-ch.ori.Y]
}

func (ch *Chart[T]) Put(x, y int, v T) {
	ch.data[y-ch.ori.Y][x-ch.ori.X] = v
}

func (ch *Chart[T]) Each(itFn func(x, y int, v T)) {
	x1, y1 := ch.ori.X, ch.ori.Y
	x2 := x1 + ch.width
	y2 := y1 + ch.height
	for i := y1; i < y2; i++ {
		for j := x1; j < x2; j++ {
			v := ch.Get(j, i)
			itFn(j, i, v)
		}
	}
}

func (ch *Chart[T]) StringF(sprintf func(v T) string) string {
	var b strings.Builder
	for i := 0; i < ch.height; i++ {
		for j := 0; j < ch.width; j++ {
			v := ch.data[i][j]
			b.WriteString(sprintf(v))
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func (row Row[T]) StringF(sprintf func(v T) string) string {
	var b strings.Builder
	l := len(row)
	for i := 0; i < l; i++ {
		b.WriteString(sprintf(row[i]))
	}
	return b.String()
}
