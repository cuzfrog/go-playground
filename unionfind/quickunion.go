package unionfind

type quUnions struct {
	baseUnions
}

func newQuUnions(n int) unions{
	u := baseUnions{n, genSites(n), nil, nil}

	u.findImpl = func(a int) int {
		for u.sites[a] != a {
			a = u.sites[a]
		}
		return a
	}

	// find searches through links and returns the component root number
	u.unionImpl = func(a int, b int) {
		roota := u.find(a)
		rootb := u.find(b)
		if roota == rootb {
			return
		}

		u.sites[roota] = rootb
		u.comCnt--
	}
	return &quUnions{u}
}
