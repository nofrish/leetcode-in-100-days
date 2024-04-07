package _054

func spiralOrder_(matrix [][]int) []int {

	M, N := len(matrix), len(matrix[0])

	up, down, left, right := 0, M-1, 0, N-1
	result := make([]int, 0, M*N)
	for {
		for x, y := up, left; y <= right; y++ {
			result = append(result, matrix[x][y])
		}
		if up++; up > down {
			break
		}

		for x, y := up, right; x <= down; x++ {
			result = append(result, matrix[x][y])
		}
		if right--; right < left {
			break
		}

		for x, y := down, right; y >= left; y-- {
			result = append(result, matrix[x][y])
		}
		if down--; down < up {
			break
		}

		for x, y := down, left; x >= up; x-- {
			result = append(result, matrix[x][y])
		}
		if left++; left > right {
			break
		}
	}
	return result
}
