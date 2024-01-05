package _977

func sortedSquares(nums []int) []int {

	result, idx := make([]int, len(nums)), len(nums)-1

	for left, right := 0, len(nums)-1; left <= right; idx-- {
		ls := nums[left] * nums[left]
		rs := nums[right] * nums[right]
		if ls < rs {
			result[idx] = rs
			right--
		} else {
			result[idx] = ls
			left++
		}
	}
	return result
}
