package _059

func generateMatrix(n int) [][]int {

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	const (
		Right = iota
		Down
		Left
		Up
	)

	NextPosition := func(x, y int, direction int) (int, int) {
		switch direction {
		case Right:
			return x, y + 1
		case Down:
			return x + 1, y
		case Left:
			return x, y - 1
		case Up:
			return x - 1, y
		}
		return 0, 0
	}

	NextDirection := func(direction int) int {
		switch direction {
		case Right:
			return Down
		case Down:
			return Left
		case Left:
			return Up
		case Up:
			return Right
		}
		return -1
	}

	IsValid := func(x, y int) bool {
		if x < 0 || x >= n || y < 0 || y >= n || result[x][y] != 0 {
			return false
		}
		return true
	}

	x, y, direction := 0, 0, Right
	result[x][y] = 1
	for i := 2; i <= n*n; i++ {
		if !IsValid(NextPosition(x, y, direction)) {
			direction = NextDirection(direction)
		}
		x, y = NextPosition(x, y, direction)
		result[x][y] = i
	}

	return result
}
