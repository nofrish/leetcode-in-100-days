package _494

func findTargetSumWays(nums []int, target int) (result int) {

	var dfs func(int, int)
	dfs = func(index, sum int) {
		if index == len(nums) {
			if sum == target {
				result++
			}
			return
		}
		dfs(index+1, sum+nums[index])
		dfs(index+1, sum-nums[index])
	}

	dfs(0, 0)
	return
}
