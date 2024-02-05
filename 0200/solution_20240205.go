package _200

func numIslands(grid [][]byte) (result int) {

	height := len(grid)
	width := len(grid[0])

	checked := make([][]bool, height)
	for i := range grid {
		checked[i] = make([]bool, width)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		nexts := make([][2]int, 0)
		if j < width-1 && grid[i][j+1] == '1' && checked[i][j+1] == false {
			checked[i][j+1] = true
			nexts = append(nexts, [2]int{i, j + 1})
		}
		if j > 0 && grid[i][j-1] == '1' && checked[i][j-1] == false {
			checked[i][j-1] = true
			nexts = append(nexts, [2]int{i, j - 1})
		}
		if i < height-1 && grid[i+1][j] == '1' {
			checked[i+1][j] = true
			nexts = append(nexts, [2]int{i + 1, j})
		}
		if i > 0 && grid[i-1][j] == '1' && checked[i-1][j] == false {
			checked[i-1][j] = true
			nexts = append(nexts, [2]int{i - 1, j})
		}
		for _, v := range nexts {
			dfs(v[0], v[1])
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && checked[i][j] == false {
				checked[i][j] = true
				result++
				dfs(i, j)
			}
		}
	}

	return
}
