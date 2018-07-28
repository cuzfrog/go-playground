package unionfind

type wquUnions struct {
	*abstractUnions
}

func newWquUnions(n int) unions {
	u := &abstractUnions{n, genSites(n), nil, nil}

	u.findImpl = u.findLink

	u.unionImpl = func(a int, b int) {
		rootA, dpthA := u.findDepth(a)
		rootB, dpthB := u.findDepth(b)
		if rootA != rootB {
			if dpthA <= dpthB {
				u.sites[rootA] = rootB
			} else {
				u.sites[rootB] = rootA
			}
			u.comCnt--
		}
	}

	return &wquUnions{u}
}

// find searches through links and returns the component root number and tree depth
func (u *abstractUnions) findDepth(a int) (int, depth int) {
	for u.sites[a] != a {
		a = u.sites[a]
		depth++
	}
	return a, depth
}
