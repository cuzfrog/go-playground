package searching

type dictionary interface {
	get(k int) interface{}
	put(k int, v interface{}) (interface{}, error)
	remove(k int) interface{}
	contains(k int) bool
	size() int
	iterator() chan entry
}

type entry struct {
	k int
	v interface{}
}

