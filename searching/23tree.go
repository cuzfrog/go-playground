package searching

type twoThreeTree struct {
	n    int
	root node23
}

type node23 struct {
	el    entry
	er    entry
	left  node23
	mid   node23
	right node23
}

func (*twoThreeTree) get(k int) interface{} {
	panic("implement me")
}

func (*twoThreeTree) put(k int, v interface{}) {
	panic("implement me")
}

func (*twoThreeTree) remove(k int) interface{} {
	panic("implement me")
}

func (*twoThreeTree) contains(k int) bool {
	panic("implement me")
}

func (*twoThreeTree) size() int {
	panic("implement me")
}

func (*twoThreeTree) iterator() chan entry {
	panic("implement me")
}

