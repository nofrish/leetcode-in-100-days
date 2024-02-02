package _509

func fib(n int) int {
	results := make([]int, n+3)
	results[0] = 0
	results[1] = 1
	for i := 2; i <= n; i++ {
		results[i] = results[i-1] + results[i-2]
	}
	return results[n]
}
