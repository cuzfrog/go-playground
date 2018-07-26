package unionfind

import (
	"testing"
	"fmt"
)

func TestUnions(t *testing.T) {
	u := newUnions(100)
	u.union(3, 6)
	u.union(8, 11)
	fmt.Println(u.count())
}