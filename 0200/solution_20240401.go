package _200

func numIslands__(grid [][]byte) int {

	var dfs func(x, y int)
	dfs = func(x, y int) { // 对于一个岛屿点，标记所有与它相连的点
		if x < 0 || y < 0 || x > len(grid)-1 || y > len(grid[0])-1 {
			return
		}
		if grid[x][y] != '1' {
			return
		}
		grid[x][y] = 'X'
		dfs(x+1, y)
		dfs(x, y+1)
		dfs(x-1, y)
		dfs(x, y-1)
	}

	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				result += 1
				dfs(i, j)
			}
		}
	}
	return result
}
