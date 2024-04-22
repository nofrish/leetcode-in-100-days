package _048

func rotate_(matrix [][]int) {

	N := len(matrix)

	for x := 0; x < N/2; x++ {
		for y := 0; y < N; y++ {
			matrix[x][y], matrix[N-1-x][y] = matrix[N-1-x][y], matrix[x][y]
		}
	}

	for x := 0; x < N; x++ {
		for y := x + 1; y < N; y++ {
			matrix[x][y], matrix[y][x] = matrix[y][x], matrix[x][y]
		}
	}
}
