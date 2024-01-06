package _238

func productExceptSelf(nums []int) []int {

	forwords := make([]int, len(nums)+1)
	forwords[0] = 1
	for i := 1; i < len(forwords); i++ {
		forwords[i] = forwords[i-1] * nums[i-1]
	}

	backwords := make([]int, len(nums)+1)
	backwords[len(backwords)-1] = 1
	for i := len(backwords) - 2; i >= 0; i-- {
		backwords[i] = backwords[i+1] * nums[i]
	}

	result := make([]int, len(nums))
	for i := range nums {
		result[i] = forwords[i] * backwords[i+1]
	}

	return result
}
