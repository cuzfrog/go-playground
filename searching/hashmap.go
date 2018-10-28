package searching

type hashTable struct {
	table []linkedNode
	cnt   int
}

type linkedNode struct {
	k    int
	v    interface{}
	next linkedNode
}


func (*hashTable) get(k int) interface{} {
	panic("implement me")
}

func (*hashTable) put(k int, v interface{}) (interface{}, error) {
	panic("implement me")
}

func (*hashTable) remove(k int) interface{} {
	panic("implement me")
}

func (*hashTable) contains(k int) bool {
	panic("implement me")
}

func (*hashTable) size() int {
	panic("implement me")
}
