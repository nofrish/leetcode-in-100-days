package _977

func sortedSquares__(nums []int) []int {

	result := make([]int, len(nums))
	i, j, cur := 0, len(nums)-1, len(result)-1
	for ; i <= j; cur-- {
		s1 := nums[i] * nums[i]
		s2 := nums[j] * nums[j]
		if s1 >= s2 {
			result[cur] = s1
			i++
		} else {
			result[cur] = s2
			j--
		}
	}
	return result
}
