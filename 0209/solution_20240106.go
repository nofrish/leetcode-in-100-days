package _209

import "math"

func minSubArrayLen_(target int, nums []int) int {
	slow, fast, sum, min := 0, -1, 0, math.MaxInt
	for fast++; fast < len(nums); fast++ {
		sum += nums[fast]
		for ; sum >= target; slow++ {
			if min > fast-slow+1 {
				min = fast - slow + 1
			}
			sum -= nums[slow]
		}
	}

	if min == math.MaxInt {
		min = 0
	}
	return min
}
