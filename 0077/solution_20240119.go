package _077

func combine(n int, k int) (result [][]int) {

	var dfs func(prefix []int, remains []int)
	dfs = func(prefix []int, remains []int) {
		if len(prefix)+len(remains) < k {
			return
		}
		if len(prefix) == k {
			result = append(result, prefix)
			return
		}
		for i, v := range remains {
			dfs(append(append([]int{}, prefix...), v), append([]int{}, remains[i+1:]...))
		}
	}

	remains := make([]int, n)
	for i := 0; i < n; i++ {
		remains[i] = i + 1
	}

	dfs([]int{}, remains)
	return
}
