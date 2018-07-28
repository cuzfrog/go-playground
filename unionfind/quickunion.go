package unionfind

type quUnions struct {
	comCnt int
	sites []int
}

func newQuUnions(n int) unions{
	return &quUnions{n, genSites(n)}
}

func (u *quUnions) union(a, b int) {
	roota := u.find(a)
	rootb := u.find(b)
	if roota == rootb {
		return
	}

	u.sites[roota] = rootb
	u.comCnt--
}

// find searches through links and returns the component root number
func (u *quUnions) find(a int) int {
	for u.sites[a] != a {
		a = u.sites[a]
	}
	return a
}

func (u *quUnions) connected(a, b int) bool {
	return u.find(a) == u.find(b)
}

func (u *quUnions) count() int {
	return u.comCnt
}
