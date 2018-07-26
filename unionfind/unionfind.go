package unionfind


type unions struct {
	comCnt int
	sites  []int
}

//new creates a new unions with given 'n' site capacity
func newUnions(n int) unions {
	sites := make([]int, n, n)
	for i := range sites {
		sites[i] = i
	}
	return unions{n, sites}
}

//union connects two sites
func (u *unions) union(a, b int) {
	ida := u.sites[a]
	idb := u.sites[b]
	if ida != idb {
		u.sites[b] = ida
		u.comCnt--
	}
}

//find returns component id for a given site
func (u *unions) find(a int) int {
	return u.sites[a]
}

//connected checks if two sites are in the same union
func (u *unions) connected(a, b int) bool {
	return u.sites[a] == u.sites[b]
}

//count returns the number of components
func (u *unions) count() int {
	return u.comCnt
}
