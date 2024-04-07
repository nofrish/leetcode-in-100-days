package _048

func rotate(matrix [][]int) {

	N := len(matrix)

	// 先上下翻转
	for i := 0; i < N/2; i++ {
		for j := 0; j < N; j++ {
			matrix[i][j], matrix[N-i-1][j] = matrix[N-i-1][j], matrix[i][j]
		}
	}
	// 再对角线翻转（左上到右下）
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	return
}
