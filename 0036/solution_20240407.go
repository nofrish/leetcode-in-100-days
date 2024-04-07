package _036

func isValidSudoku(board [][]byte) bool {

	const N = 9
	rows, columns, squares := [N][N]int{}, [N][N]int{}, [N][N]int{}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == '.' {
				continue
			}
			index := int(board[i][j] - '1')
			rows[i][index] += 1
			columns[j][index] += 1
			squares[i/3*3+j/3][index] += 1
			if rows[i][index] > 1 || columns[j][index] > 1 || squares[i/3*3+j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
