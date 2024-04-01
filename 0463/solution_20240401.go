package _463

func islandPerimeter__(grid [][]int) int {

	inArea := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
	}

	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 1 {
				continue
			}
			for _, move := range [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				if !inArea(i+move[0], j+move[1]) || grid[i+move[0]][j+move[1]] == 0 {
					result++
				}
			}
		}
	}
	return result
}
