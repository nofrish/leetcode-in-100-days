package _073

func setZeroes(matrix [][]int) {

	rows, columns := make(map[int]struct{}), make(map[int]struct{})
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = struct{}{}
				columns[j] = struct{}{}
			}
		}
	}

	for row := range rows {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[row][j] = 0
		}
	}

	for column := range columns {
		for i := 0; i < len(matrix); i++ {
			matrix[i][column] = 0
		}
	}

}
