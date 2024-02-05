package _463

func islandPerimeter(grid [][]int) (result int) {

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if !inArea(grid, i, j) || grid[i][j] == 0 { // 遇到岛屿边界
			result++
			return
		}
		if grid[i][j] == -1 { // 遇到检查过的陆地
			return
		}

		grid[i][j] = -1

		dfs(i, j+1)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i-1, j)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				dfs(i, j)
				return
			}
		}
	}
	return
}

func inArea(grid [][]int, i, j int) bool {
	return 0 <= i && 0 <= j && i < len(grid) && j < len(grid[0])
}
