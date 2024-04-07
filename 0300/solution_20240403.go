package _300

func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}

	result := 0
	for i := 0; i < len(dp); i++ {
		result = max(result, dp[i])
	}
	return result
}
