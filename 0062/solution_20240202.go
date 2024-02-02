package _062

func uniquePaths(m int, n int) int {

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	var dfs func(m, n int) int
	dfs = func(m, n int) int {
		if m == 0 || n == 0 {
			return 0
		}
		if m == 1 || n == 1 {
			return 1
		}
		if dp[m][n] != 0 {
			return dp[m][n]
		}
		dp[m][n] = dfs(m-1, n) + dfs(m, n-1)
		return dp[m][n]
	}

	return dfs(m, n)
}
