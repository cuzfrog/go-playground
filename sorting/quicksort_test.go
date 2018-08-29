package sorting

import "testing"

func TestPartition(t *testing.T) {
	for n := 0; n < 100; n++ {
		a := genElems(16)
		t.Log("Original:", a)
		j := partition(a)
		l := a[:j]
		r := a[j+1:]

		for il, el := range l {
			for ir, er := range r {
				if el > er {
					t.Logf("Partition left: %v", l)
					t.Logf("Partition right:%v", r)
					t.Logf("Element %v at %d in left is greater than elem %v at %d in right", el, il, er, ir)
					t.Logf("Key elem is %v", a[j])
					t.FailNow()
				}
			}
		}
	}
}
