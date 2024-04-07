package _221

func maximalSquare(matrix [][]byte) int {

	square := func(x, y int) int {
		area := 1
		for delta := 1; ; delta++ {
			if x+delta >= len(matrix) || y+delta >= len(matrix[0]) {
				return area
			}
			for i, j := x+delta, y; j < y+delta; j++ {
				if matrix[i][j] != '1' {
					return area
				}
			}
			for i, j := x, y+delta; i <= x+delta; i++ {
				if matrix[i][j] != '1' {
					return area
				}
			}
			area = (1 + delta) * (1 + delta)
		}
		return area
	}

	maxArea := 0
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[0]); y++ {
			if matrix[x][y] == '1' {
				area := square(x, y)
				maxArea = max(maxArea, area)
			}
		}
	}
	return maxArea
}
