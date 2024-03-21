package _977

func sortedSquares___(nums []int) []int {

	result := make([]int, len(nums))

	i, j, cur := 0, len(nums)-1, len(result)-1
	for ; i <= j; cur-- {
		ii := nums[i] * nums[i]
		jj := nums[j] * nums[j]
		if ii >= jj {
			result[cur] = ii
			i++
		} else {
			result[cur] = jj
			j--
		}
	}

	return result
}
