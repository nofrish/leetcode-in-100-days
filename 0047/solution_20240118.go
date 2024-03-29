package _047

func permuteUnique(nums []int) (result [][]int) {

	var dfs func(prefix []int, options []int)
	dfs = func(prefix []int, options []int) {
		if len(options) == 0 {
			result = append(result, prefix)
			return
		}
		used := make(map[int]bool)
		for i, v := range options {
			if !used[v] {
				dfs(append(append([]int{}, prefix...), v), append(append([]int{}, options[0:i]...), options[i+1:]...))
				used[v] = true
			}
		}
	}

	dfs([]int{}, nums)
	return
}
