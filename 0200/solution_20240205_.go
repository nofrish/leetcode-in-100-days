package _200

func numIslands_(grid [][]byte) (result int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				result++
				dfs(grid, i, j)
			}
		}
	}
	return
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[0])-1 {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = 'x'

	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
}
