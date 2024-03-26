package _123

func maxProfit(prices []int) int {

	N := len(prices)

	dp := make([][4]int, N)
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = -prices[0]
	dp[0][3] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]-prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]+prices[i])
	}

	return dp[N-1][3]
}
