package _059

func generateMatrix_(n int) [][]int {

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	i, up, down, left, right := 1, 0, n-1, 0, n-1
	for {
		for x, y := up, left; y <= right; y++ {
			matrix[x][y] = i
			i++
		}
		if up++; up > down {
			break
		}

		for x, y := up, right; x <= down; x++ {
			matrix[x][y] = i
			i++
		}
		if right--; right < left {
			break
		}

		for x, y := down, right; y >= left; y-- {
			matrix[x][y] = i
			i++
		}
		if down--; down < up {
			break
		}

		for x, y := down, left; x >= up; x-- {
			matrix[x][y] = i
			i++
		}
		if left++; left > right {
			break
		}
	}

	return matrix
}
