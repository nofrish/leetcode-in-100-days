package _209

import "math"

func minSubArrayLen(target int, nums []int) int {

	slow, fast, sum, min := 0, 0, nums[0], math.MaxInt
	for {
		if sum < target {
			fast++
			if fast == len(nums) {
				break
			}
			sum += nums[fast]
		} else {
			if min > fast-slow+1 {
				min = fast - slow + 1
			}
			sum -= nums[slow]
			slow++
		}
	}

	if min == math.MaxInt {
		min = 0
	}
	return min
}
