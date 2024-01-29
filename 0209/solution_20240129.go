package _209

import "math"

func minSubArrayLen___(target int, nums []int) int {

	slow, fast, minLen, sum := 0, -1, math.MaxInt, 0
	for fast++; fast < len(nums); fast++ {
		sum += nums[fast]
		for ; sum >= target; slow++ {
			if minLen > fast-slow+1 {
				minLen = fast - slow + 1
			}
			sum -= nums[slow]
		}
	}

	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}
