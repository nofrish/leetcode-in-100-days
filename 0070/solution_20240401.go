package _070

func climbStairs_(n int) int {
	dp := []int{1, 2}
	for i := 3; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}
