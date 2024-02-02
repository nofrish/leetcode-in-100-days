package _070

func climbStairs(n int) int {
	results := make([]int, n+3)
	results[1] = 1
	results[2] = 2
	for i := 3; i <= n; i++ {
		results[i] = results[i-1] + results[i-2]
	}
	return results[n]
}
