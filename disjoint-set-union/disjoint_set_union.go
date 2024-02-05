package disjoint_set_union

type DisjointSet struct {
	parents []int
}

func NewDisjointSet(size int) *DisjointSet {
	ds := &DisjointSet{parents: make([]int, size)}
	for i := range ds.parents {
		ds.parents[i] = i
	}
	return ds
}

func (ds *DisjointSet) Find(x int) int {
	if ds.parents[x] != x {
		ds.parents[x] = ds.Find(ds.parents[x]) // with path compression
	}
	return ds.parents[x]
}

func (ds *DisjointSet) Union(x, y int) {
	ds.parents[ds.Find(x)] = ds.Find(y)
}
