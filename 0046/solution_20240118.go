package _046

func permute(nums []int) (result [][]int) {

	var dfs func(prefix []int, remains []int)
	dfs = func(prefix []int, remains []int) {
		if len(remains) == 0 {
			result = append(result, prefix)
			return
		}
		for i, v := range remains {
			dfs(append(append([]int{}, prefix...), v), append(append([]int{}, remains[0:i]...), remains[i+1:]...))
		}
	}

	dfs([]int{}, nums)
	return
}
