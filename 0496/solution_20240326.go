package _496

func nextGreaterElement_(nums1 []int, nums2 []int) []int {

	// positions用于记录nums1中每一个数字对应的下标
	// 我的思路是，对nums2中的每一个数字求它的下一个更大元素
	// 当发现该数字也在positions中时，更新answer对应的位置
	positions := make(map[int]int, len(nums1))
	for i, v := range nums1 {
		positions[v] = i
	}

	// 将answer的每一个位置都初始化为-1
	answer := make([]int, len(nums1))
	for i := range answer {
		answer[i] = -1
	}

	stack := make([]int, 0, len(nums2))
	for _, v := range nums2 {
		for len(stack) != 0 && v > stack[len(stack)-1] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if idx, ok := positions[pop]; ok {
				answer[idx] = v
			}
		}
		stack = append(stack, v) // 注意：本题只要往单调栈中写入值就好，不需要写入下标
	}

	return answer
}
