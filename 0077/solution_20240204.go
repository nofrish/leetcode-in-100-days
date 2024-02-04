package _077

func combine(n int, k int) (result [][]int) {

	var dfs func(prefix []int, options []int)
	dfs = func(prefix []int, options []int) {
		if len(prefix) == k {
			result = append(result, prefix)
			return
		}
		for i, v := range options {
			if i+k > n {
				return
			}
			dfs(append(append([]int{}, prefix...), v), append([]int{}, options[i+1:]...))
		}
	}

	options := make([]int, n)
	for i := 0; i < n; i++ {
		options[i] = i + 1
	}
	dfs([]int{}, options)

	return
}
