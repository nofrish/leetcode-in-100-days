package _047

func permuteUnique(nums []int) (result [][]int) {

	var dfs func(prefix []int, remains []int)
	dfs = func(prefix []int, remains []int) {
		if len(remains) == 0 {
			result = append(result, prefix)
			return
		}
		used := make(map[int]bool)
		for i, v := range remains {
			if !used[v] {
				dfs(append(append([]int{}, prefix...), v), append(append([]int{}, remains[0:i]...), remains[i+1:]...))
				used[v] = true
			}
		}
	}

	dfs([]int{}, nums)
	return
}
