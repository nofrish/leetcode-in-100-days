package _121

func maxProfit_(prices []int) int {

	min, result := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if result < prices[i]-min {
			result = prices[i] - min
		}
		if min > prices[i] {
			min = prices[i]
		}
	}

	return result
}

func maxProfit(prices []int) int {

	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0])
	}

	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
