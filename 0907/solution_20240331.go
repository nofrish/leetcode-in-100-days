package _907

func sumSubarrayMins(nums []int) int {
	stack := make([]int, 0, len(nums))
	left, right := make([]int, len(nums)), make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] <= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = stack[0:0]
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[i] < nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = len(nums)
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i] * (i - left[i]) * (right[i] - i)
	}
	return sum % (1e9 + 7)
}
