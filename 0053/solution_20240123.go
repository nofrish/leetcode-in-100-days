package _053

import "math"

func maxSubArray(nums []int) int {
	max := math.MinInt
	for i, sum := 0, 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
