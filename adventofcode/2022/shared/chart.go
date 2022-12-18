package shared

type Chart[T comparable] struct {
	data   [][]T
	ori    Coord
	width  int
	height int
}

func NewChart[T comparable](oriX, oriY int, w, h int) *Chart[T] {
	data := make([][]T, h)
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

func (ch *Chart[T]) Put(x, y int, v T) {
	ch.data[y-ch.ori.Y][x-ch.ori.X] = v
}
