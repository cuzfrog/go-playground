package searching

import (
	"go-playground/utils"
	"fmt"
)

type hashTable struct {
	table []*linkedNode
	cnt   int
}

type linkedNode struct {
	k    int
	v    interface{}
	next *linkedNode
}

func (t *hashTable) get(k int) interface{} {
	n := locateNode(t, k)
	if n != nil {
		return n.v
	} else {
		return nil
	}
}

func (t *hashTable) put(k int, v interface{}) (old interface{}, err error) {
	if v == nil {
		return nil, fmt.Errorf("v cannot be nil")
	}
	if t.cnt >= cap(t.table) {
		increaseCapacity(t)
	}
	h := utils.Hash(k, cap(t.table))
	n := t.table[h]
	if n == nil {
		t.table[h] = &linkedNode{k, v, nil}
	} else {
		for n != nil {
			if n.k == k {
				old, n.v = n.v, v
				break
			}
			if n.next == nil {
				n.next = &linkedNode{k, v, nil}
				break
			}
			n = n.next
		}
	}
	if old == nil {
		t.cnt++
	}
	return old, err
}

func (t *hashTable) remove(k int) interface{} {
	panic("implement me")
}

func (t *hashTable) contains(k int) bool {
	return locateNode(t, k) != nil
}

func (t *hashTable) size() int {
	return t.cnt
}

func locateNode(t *hashTable, k int) *linkedNode {
	h := utils.Hash(k, cap(t.table))
	return t.table[h]
}

func increaseCapacity(t *hashTable) {
	//todo
}
