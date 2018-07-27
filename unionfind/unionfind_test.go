package unionfind

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUnions(t *testing.T) {
	Assert := assert.New(t)
	impls := []unions{
		newQfUnions(100),
	}

	for _, u := range impls {
		u.union(3, 6)
		u.union(8, 11)
		u.union(3, 9)
		u.union(67, 9)
		u.union(6, 9)
		Assert.Equal(96, u.count(), "Component num is wrong")
		Assert.True(u.connected(67, 6))
		Assert.False(u.connected(8, 9))
	}
}
