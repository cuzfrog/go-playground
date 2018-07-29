package unionfind

// unions with quick find implementation
type qfUnions *delegateUnions

func newQfUnions(n int) qfUnions {
	u := &delegateUnions{n, genSites(n), nil, nil}

	u.findImpl = func(a int) int {
		return u.sites[a]
	}

	u.unionImpl = func(a int, b int) {
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

	return u
}
