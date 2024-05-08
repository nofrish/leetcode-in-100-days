package _240

func searchMatrix_(matrix [][]int, target int) bool {

	M, N := len(matrix), len(matrix[0])

	for i := 0; i < M; i++ {
		if matrix[i][N-1] < target {
			// 如果某一行的最后一个元素小于target，说明整行都小，跳过
			continue
		}
		if matrix[i][0] > target {
			// 如果某一行的第一个元素大于target，说明整行都大，未来的值更大，退出循环
			break
		}

		lo, hi := 0, N-1
		for lo <= hi {
			mid := (lo + hi) >> 1
			if matrix[i][mid] < target {
				lo = mid + 1
			} else if matrix[i][mid] > target {
				hi = mid - 1
			} else {
				return true
			}
		}
	}
	return false
}
