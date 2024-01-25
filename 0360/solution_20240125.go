package _360

func sortTransformedArray(nums []int, a int, b int, c int) []int {

	result := make([]int, len(nums))

	cur := 0
	if a > 0 {
		cur = len(result) - 1
	}

	for i, j := 0, len(nums)-1; i <= j; {
		r1 := a*nums[i]*nums[i] + b*nums[i] + c
		r2 := a*nums[j]*nums[j] + b*nums[j] + c

		if a < 0 {
			if r1 < r2 {
				result[cur] = r1
				i++
			} else {
				result[cur] = r2
				j--
			}
			cur++
		} else {
			if r1 > r2 {
				result[cur] = r1
				i++
			} else {
				result[cur] = r2
				j--
			}
			cur--
		}
	}

	return result
}
