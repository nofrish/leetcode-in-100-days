package _122

func maxProfit_(prices []int) int {

	N := len(prices)

	dp := make([][2]int, N)
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for i := 1; i < N; i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][0] = max(dp[i-1][0], dp[i][1]-prices[i]) // 注意要用今天不持有股票的利润来计算
	}

	return dp[N-1][1]
}
