package _042

func trap_(height []int) int {
	stack := []int{-1}
	left, right := make([]int, len(height)), make([]int, len(height))

	for i := 0; i < len(height); i++ {
		left[i], right[i] = -1, -1
		for len(stack) > 1 && height[i] >= height[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			right[pop] = i
		}
		left[i] = stack[len(stack)-1]
		stack = append(stack, i)
	}

	sum := 0
	for i := 0; i < len(height); i++ {
		if left[i] != -1 && right[i] != -1 {
			sum += (right[i] - left[i] - 1) * min(height[left[i]]-height[i], height[right[i]]-height[i])
		}
	}
	return sum
}
