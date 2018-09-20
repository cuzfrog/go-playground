package searching

import "fmt"

type rbtree struct {
	count int
	root  *rbnode
}

func newRbtreeDic() dictionary {
	return &rbtree{}
}

func (t *rbtree) get(k int) interface{} {
	return t.root.get(k)
}

func (t *rbtree) put(k int, v interface{}) (old interface{}, err error) {
	if v == nil {
		err = fmt.Errorf("v cannot be nil")
	} else {
		t.root, old = t.root.insert(k, v)
		if old == nil {
			t.count++
		}
	}
	return
}

func (t *rbtree) remove(k int) (old interface{}) {
	t.root, old = t.root.remove(k)
	if old != nil {
		t.count--
	}
	return
}

func (t *rbtree) contains(k int) bool {
	return t.get(k) != nil
}

func (t *rbtree) size() int {
	return t.count;
}

func (*rbtree) iterator() chan entry {
	panic("implement me")
}
