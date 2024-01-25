package _977

func sortedSquares_(nums []int) []int {

	result := make([]int, len(nums))
	for i, j, c := 0, len(nums)-1, len(nums)-1; i <= j; c-- {
		r1 := nums[i] * nums[i]
		r2 := nums[j] * nums[j]
		if r1 > r2 {
			result[c] = r1
			i++
		} else {
			result[c] = r2
			j--
		}
	}
	return result
}
