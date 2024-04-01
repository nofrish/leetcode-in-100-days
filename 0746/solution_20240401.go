package _746

func minCostClimbingStairs_(cost []int) int {
	dp := []int{cost[0], cost[1]}
	for i := 2; i < len(cost); i++ {
		dp = append(dp, cost[i]+min(dp[i-1], dp[i-2]))
	}
	return min(dp[len(dp)-1], dp[len(dp)-2])
}
