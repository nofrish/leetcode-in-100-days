package _016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	result := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			if left > i+1 && nums[left] == nums[left-1] { // deduplicate left
				left++
				continue
			}
			if right < len(nums)-1 && nums[right] == nums[right+1] { // deduplicate right
				right--
				continue
			}

			sum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(result-target)) > math.Abs(float64(sum-target)) {
				result = sum
			}
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return sum
			}
		}
	}

	return result
}
