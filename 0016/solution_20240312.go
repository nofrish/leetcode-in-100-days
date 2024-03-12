package _016

import (
	"math"
	"sort"
)

func threeSumClosest__(nums []int, target int) int {

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	sort.Ints(nums)

	result := math.MaxInt
	for i := range nums {
		low, high := i+1, len(nums)-1
		for low < high {
			sum := nums[i] + nums[low] + nums[high]
			if abs(sum-target) < abs(result-target) {
				result = sum
			}
			if sum > target {
				high--
			} else if sum < target {
				low++
			} else {
				return sum
			}
		}
	}
	return result
}
