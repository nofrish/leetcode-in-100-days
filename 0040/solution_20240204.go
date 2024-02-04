package _040

import "sort"

func combinationSum2(nums []int, target int) (result [][]int) {

	sort.Ints(nums)

	var dfs func(prefix []int, options []int, k int)
	dfs = func(prefix []int, options []int, k int) {
		if len(prefix) > k {
			return
		}
		sum := 0
		for _, v := range prefix {
			sum += v
		}
		if sum > target {
			return
		}
		if len(prefix) == k && sum == target {
			result = append(result, prefix)
		}

		for i := range options {
			if i > 0 && options[i] == options[i-1] {
				continue
			}
			dfs(append(append([]int{}, prefix...), options[i]), append([]int{}, options[i+1:]...), k)
		}
	}

	for k := 1; k <= len(nums); k++ {
		dfs([]int{}, nums, k)
	}

	return
}
