package _213

func rob(nums []int) (result int) {

	rob_ := func(nums []int) int {
		pre, cur := 0, nums[0]
		for i := 1; i < len(nums); i++ {
			cur, pre = max(cur, pre+nums[i]), cur
		}
		return cur
	}

	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob_(nums[1:]), rob_(nums[:len(nums)-1]))
}
