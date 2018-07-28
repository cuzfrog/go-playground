package unionfind

type unions interface {
	//union connects two sites
	union(a, b int)
	//connected checks if two sites are in the same union
	connected(a, b int) bool
	//count returns the number of components
	count() int

	find(a int) int
}

type baseUnions struct {
	comCnt    int
	sites     []int
	findImpl  func(int) int
	unionImpl func(int, int)
}

func (u *baseUnions) union(a, b int) {
	u.unionImpl(a, b)
}

func (u *baseUnions) find(a int) int {
	return u.findImpl(a)
}

func (u *baseUnions) connected(a, b int) bool {
	return u.findImpl(a) == u.findImpl(b)
}

func (u *baseUnions) count() int {
	return u.comCnt
}

/* --- Utility functions --- */

func genSites(n int) []int {
	sites := make([]int, n, n)
	for i := range sites {
		sites[i] = i
	}
	return sites
}
