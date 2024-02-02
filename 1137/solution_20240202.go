package _137

func tribonacci(n int) int {
	results := []int{0, 1, 1}
	for i := 3; i <= n; i++ {
		results = append(results, results[i-1]+results[i-2]+results[i-3])
	}
	return results[n]
}
