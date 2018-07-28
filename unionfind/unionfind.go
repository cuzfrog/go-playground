package unionfind

type unions interface {
	// union connects two sites
	union(a, b int)
	// connected checks if two sites are in the same union
	connected(a, b int) bool
	// count returns the number of components
	count() int
	// find returns component id
	find(a int) int
}

type abstractUnions struct {
	comCnt    int
	sites     []int
	findImpl  func(int) int
	unionImpl func(int, int)
}

func (u *abstractUnions) union(a, b int) {
	u.unionImpl(a, b)
}

func (u *abstractUnions) find(a int) int {
	return u.findImpl(a)
}

func (u *abstractUnions) connected(a, b int) bool {
	return u.findImpl(a) == u.findImpl(b)
}

func (u *abstractUnions) count() int {
	return u.comCnt
}

func (u *abstractUnions) findLink(a int) int {
	for u.sites[a] != a {
		a = u.sites[a]
	}
	return a
}

/* --- Utility functions --- */

func genSites(n int) []int {
	sites := make([]int, n, n)
	for i := range sites {
		sites[i] = i
	}
	return sites
}
