package _047

import "sort"

func permuteUnique__(nums []int) (result [][]int) {

	sort.Ints(nums)

	var dfs func(prefix []int, options []int)
	dfs = func(prefix []int, options []int) {
		if len(options) == 0 {
			result = append(result, prefix)
			return
		}
		last := -17
		for i, v := range options {
			if last == -17 || v != last {
				dfs(append(append([]int{}, prefix...), v), append(append([]int{}, options[0:i]...), options[i+1:]...))
				last = v
			}
		}
	}

	dfs([]int{}, nums)
	return
}
