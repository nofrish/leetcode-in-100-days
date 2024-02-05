package _463

func islandPerimeter_(grid [][]int) (result int) {

	inArea := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			for _, v := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				if !inArea(i+v[0], j+v[1]) || grid[i+v[0]][j+v[1]] == 0 {
					result++
				}
			}
		}
	}
	return
}
