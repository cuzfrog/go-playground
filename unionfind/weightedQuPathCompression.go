package unionfind

// weighted quick union with path compression
type wqupcUnions struct {
	comCnt int
	sites  []int
}

func newWqupcUnions(n int) unions {
	return &wqupcUnions{n, genSites(n)}
}

func (u *wqupcUnions) union(a, b int) {
	rootA, dpthA := u.findDepth(a)
	rootB, dpthB := u.findDepth(b)
	if rootA == rootB {
		return
	}
	if dpthA <= dpthB {
		u.sites[rootA] = rootB
	} else {
		u.sites[rootB] = rootA
	}
	u.comCnt--
}

func (u *wqupcUnions) find(a int) int {
	for u.sites[a] != a {
		a = u.sites[a]
	}
	return a
}

// find searches through links and returns the component root number and tree depth
// try to shorten path on every tracing
func (u *wqupcUnions) findDepth(a int) (int, depth int) {
	lastIdx := a
	for u.sites[a] != a {
		a = u.sites[a]
		depth++
		u.sites[lastIdx] = u.sites[a]
		lastIdx = a
	}
	return a, depth
}

func (u *wqupcUnions) connected(a, b int) bool {
	return u.find(a) == u.find(b)
}

func (u *wqupcUnions) count() int {
	return u.comCnt
}
