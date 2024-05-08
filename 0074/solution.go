package _074

func _searchMatrix(matrix [][]int, target int) bool {

	M, N := len(matrix), len(matrix[0])

	lo, hi := 0, M*N-1
	for lo <= hi {
		mid := (lo + hi) >> 1
		x, y := mid/N, mid%N
		if matrix[x][y] < target {
			lo = mid + 1
		} else if matrix[x][y] > target {
			hi = mid - 1
		} else {
			return true
		}
	}
	return false
}
