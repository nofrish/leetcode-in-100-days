package _074

func searchMatrix(matrix [][]int, target int) bool {

	// 依据矩阵第一列寻找最后一个小于目标值的行，这一行的右侧才有可能出现目标值
	// 这是二分查找解决的扩展问题
	low, high, row := 0, len(matrix)-1, -1
	for low <= high {
		row = low + (high-low)/2
		if matrix[row][0] > target {
			high = row - 1
		} else {
			if row == len(matrix)-1 || matrix[row+1][0] > target {
				break
			}
			low = row + 1
		}
	}
	if row == -1 {
		return false
	}

	// 在对应的行执行标准二分查找
	low, high = 0, len(matrix[0])-1
	for low <= high {
		mid := low + (high-low)/2
		if matrix[row][mid] > target {
			high = mid - 1
		} else if matrix[row][mid] < target {
			low = mid + 1
		} else {
			return true
		}
	}
	return false
}
