package _496

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	positions := make(map[int]int, 0)
	nexts := make([]int, len(nums2))

	stack := make([]int, 0)
	for i, v := range nums2 {
		positions[v] = i
		for len(stack) > 0 && v > nums2[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			nexts[top] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for _, v := range stack {
		nexts[v] = -1
	}

	result := make([]int, len(nums1))
	for i, v := range nums1 {
		result[i] = nexts[positions[v]]
	}
	return result
}
