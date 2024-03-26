package _503

func nextGreaterElements_(nums []int) []int {

	N := len(nums)

	answer := make([]int, N)
	for i := range nums {
		answer[i] = -1
	}

	stack := make([]int, 0, 2*N)
	for i := 0; i < 2*N; i++ {
		j := i % N
		for len(stack) != 0 && nums[j] > nums[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[idx] = nums[j]
		}
		stack = append(stack, j)
	}

	return answer
}
