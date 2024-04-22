package _054

func _spiralOrder(matrix [][]int) (result []int) {

	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
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
