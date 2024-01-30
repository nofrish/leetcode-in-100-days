package _503

func nextGreaterElements(nums []int) []int {

	result := make([]int, len(nums))

	stack := make([]int, 0)
	for i, v := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < v {
			top := stack[len(stack)-1]
			result[top] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for _, v := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < v {
			top := stack[len(stack)-1]
			result[top] = v
			stack = stack[:len(stack)-1]
		}
	}
	for _, v := range stack {
		result[v] = -1
	}

	return result
}
