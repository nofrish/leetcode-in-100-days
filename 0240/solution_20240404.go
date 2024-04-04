package _240

func searchMatrix(matrix [][]int, target int) bool {
	x, y := len(matrix)-1, 0
	for x >= 0 && y < len(matrix[0]) {
		if matrix[x][y] < target {
			y++
		} else if matrix[x][y] > target {
			x--
		} else {
			return true
		}
	}
	return false
}
