package unionfind

// weighted quick union with path compression
type wqupcUnions struct {
	*abstractUnions
}

func newWqupcUnions(n int) unions {
	u := &abstractUnions{n, genSites(n), nil, nil}

	u.findImpl = func(a int) int {
		for u.sites[a] != a {
			a = u.sites[a]
		}
		return a
	}

	u.unionImpl = func(a int, b int) {
		rootA, dpthA := u.findDepthPathCompression(a)
		rootB, dpthB := u.findDepthPathCompression(b)
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

	return &wqupcUnions{u}
}

// find searches through links and returns the component root number and tree depth
// try to shorten path on every tracing
func (u *abstractUnions) findDepthPathCompression(a int) (int, depth int) {
	lastIdx := a
	for u.sites[a] != a {
		a = u.sites[a]
		depth++
		u.sites[lastIdx] = u.sites[a]
		lastIdx = a
	}
	return a, depth
}
