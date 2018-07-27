package unionfind

// unions with quick find implementation
type qfUnions struct {
	comCnt int
	sites  []int
}

//new creates a new qfUnions with given 'n' site capacity
func newQfUnions(n int) unions {
	return &qfUnions{n, genSites(n)}
}

func (u *qfUnions) reset() {
	for i := range u.sites {
		u.sites[i] = i
	}
}

func (u *qfUnions) union(a, b int) {
	ida := u.find(a)
	idb := u.find(b)
	if ida != idb {
		for i := range u.sites {
			if u.sites[i] == ida {
				u.sites[i] = idb
			}
		}
		u.comCnt--
	}
}

func (u *qfUnions) find(a int) int {
	return u.sites[a]
}

func (u *qfUnions) connected(a, b int) bool {
	return u.find(a) == u.find(b)
}

func (u *qfUnions) count() int {
	return u.comCnt
}
