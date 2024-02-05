package _305

func numIslands2(m int, n int, positions [][]int) (result []int) {

	inArea := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}

	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	ds := NewDisjointSet(m * n)

	count := 0
	for _, pos := range positions {
		if grid[pos[0]][pos[1]] == 1 {
			result = append(result, count)
			continue
		}

		grid[pos[0]][pos[1]] = 1
		count++
		for _, v := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			x, y := pos[0]+v[0], pos[1]+v[1]
			if !inArea(x, y) || grid[x][y] == 0 {
				continue
			}
			parent1 := ds.Find(pos[0]*n + pos[1])
			parent2 := ds.Find(x*n + y)
			if parent1 != parent2 {
				ds.Union(parent1, parent2)
				count--
			}
		}
		result = append(result, count)
	}
	return
}

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
