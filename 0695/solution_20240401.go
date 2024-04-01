package _695

func maxAreaOfIsland_(grid [][]int) int {

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || y < 0 || x > len(grid)-1 || y > len(grid[0])-1 {
			return 0
		}
		if grid[x][y] != 1 {
			return 0
		}
		grid[x][y] = -1
		count := 1
		count += dfs(x+1, y)
		count += dfs(x, y+1)
		count += dfs(x-1, y)
		count += dfs(x, y-1)
		return count
	}

	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area := dfs(i, j)
				maxArea = max(maxArea, area)
			}
		}
	}
	return maxArea
}
