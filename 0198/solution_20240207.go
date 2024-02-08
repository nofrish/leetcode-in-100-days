package _198

func _rob(nums []int) int {
	pre, cur := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		cur, pre = max(cur, pre+nums[i]), cur
	}
	return cur
}

func rob(nums []int) int {

	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]

	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-1], nums[i-1]+dp[i-2])
	}
	return dp[len(dp)-1]
}

func rob_(nums []int) int {

	dp := make([][2]int, len(nums))
	dp[0] = [2]int{nums[0], 0}

	for i := 1; i < len(nums); i++ {
		dp[i] = [2]int{nums[i] + dp[i-1][1], max(dp[i-1][0], dp[i-1][1])}
	}

	last := len(nums) - 1
	return max(dp[last][0], dp[last][1])
}
