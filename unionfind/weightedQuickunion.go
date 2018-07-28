package unionfind

type wquUnions struct {
	comCnt int
	sites  []int
}

func newWquUnions(n int) unions {
	return &wquUnions{n, genSites(n)}
}

func (u *wquUnions) union(a, b int) {
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

func (u *wquUnions) find(a int) int {
	comNum, _ := u.findDepth(a)
	return comNum
}

// find searches through links and returns the component root number and tree depth
func (u *wquUnions) findDepth(a int) (int, depth int) {
	for u.sites[a] != a {
		a = u.sites[a]
		depth++
	}
	return a, depth
}

func (u *wquUnions) connected(a, b int) bool {
	return u.find(a) == u.find(b)
}

func (u *wquUnions) count() int {
	return u.comCnt
}
