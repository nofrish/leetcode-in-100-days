package _498

func findDiagonalOrder(matrix [][]int) []int {
	M, N := len(matrix), len(matrix[0])
	result := make([]int, 0, M*N)

	inArea := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[0])
	}

	moves, step := [2][2]int{{-1, 1}, {1, -1}}, 0
	x, y := 0, 0
	for {
		result = append(result, matrix[x][y])
		if x == M-1 && y == N-1 {
			break
		}
		if !inArea(x+moves[step][0], y+moves[step][1]) {
			step ^= 1
			switch {
			case x == M-1:
				y++
			case y == N-1:
				x++
			case x == 0:
				y++
			case y == 0:
				x++
			}
		} else {
			x, y = x+moves[step][0], y+moves[step][1]
		}
	}
	return result
}
