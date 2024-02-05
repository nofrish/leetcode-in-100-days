package _695

func maxAreaOfIsland(grid [][]int) int {

	max := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				result := dfs(grid, i, j)
				if max < result {
					max = result
				}
			}
		}
	}
	return max
}

func dfs(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[0])-1 {
		return 0
	}
	if grid[i][j] != 1 {
		return 0
	}

	count := 1
	grid[i][j] = -1

	count += dfs(grid, i, j+1)
	count += dfs(grid, i, j-1)
	count += dfs(grid, i+1, j)
	count += dfs(grid, i-1, j)

	return count
}
