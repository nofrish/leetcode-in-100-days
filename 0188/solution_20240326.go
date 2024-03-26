package _188

func maxProfit(k int, prices []int) int {

	N := len(prices)

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, 2*k+1)
	}
	for i := 1; i <= k; i++ {
		dp[0][2*i-1] = -prices[0]
	}

	for i := 1; i < N; i++ {
		for j := 1; j <= 2*k; j++ {
			if j%2 == 1 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			}
		}
	}

	return dp[N-1][2*k]
}
