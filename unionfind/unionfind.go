package unionfind

type unions interface {
	//union connects two sites
	union(a, b int)
	//find returns component id for a given site
	find(a int) int
	//connected checks if two sites are in the same union
	connected(a, b int) bool
	//count returns the number of components
	count() int

	reset()
}
