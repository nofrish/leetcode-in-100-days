package _042

func trap(height []int) (result int) {

	stack := make([]int, 0)
	for r, v := range height {
		for len(stack) > 0 && v > height[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]

			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			l := stack[len(stack)-1]

			area := (r - l - 1) * (min(height[l], height[r]) - height[mid])
			result += area
		}
		stack = append(stack, r)
	}
	return
}
