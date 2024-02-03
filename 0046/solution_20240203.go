package _046

func permute_(nums []int) (result [][]int) {

	var dfs func(prefix []int, options []int)
	dfs = func(prefix []int, options []int) {
		if len(options) == 0 {
			result = append(result, prefix)
			return
		}
		for i, v := range options {
			option := v
			dfs(append(prefix, option), append(append([]int{}, options[0:i]...), options[i+1:]...))
		}
	}

	dfs([]int{}, nums)
	return
}
