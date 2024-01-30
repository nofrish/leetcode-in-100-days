package _739

func dailyTemperatures(temperatures []int) []int {

	result := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			top := temperatures[stack[len(stack)-1]]
			result[top] = i - top + 1
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for _, v := range stack {
		result[v] = 0
	}

	return result
}
