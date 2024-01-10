package disjoint_set_union

type DSU struct {
	parants []int
}

func NewDSU(size int) *DSU {
	return &DSU{parants: make([]int, size)}
}

func (dsu *DSU) Find(i int) int {
	if dsu.parants[i] != i {
		dsu.parants[i] = dsu.Find(dsu.parants[i])
	}
	return dsu.parants[i]
}

func (dsu *DSU) Union(i, j int) {
	dsu.parants[i] = dsu.Find(j)
}
