package _054

func spiralOrder(matrix [][]int) (result []int) {

	m, n := len(matrix), len(matrix[0])
	up, down, left, right := 0, m-1, 0, n-1

	for {
		for x, y := up, left; y <= right; y++ {
			result = append(result, matrix[x][y])
		}
		up++
		for x, y := up, right; x <= down; x++ {
			result = append(result, matrix[x][y])
		}
		right--
		for x, y := down, right; y >= left; y-- {
			result = append(result, matrix[x][y])
		}
		down--
		for x, y := down, left; x >= up; x-- {
			result = append(result, matrix[x][y])
		}
		left++

		if len(result) >= m*n {
			result = result[:m*n]
			break
		}
	}

	return
}
