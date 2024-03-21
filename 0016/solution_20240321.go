package _016

import (
	"math"
	"sort"
)

func threeSumClosest___(nums []int, target int) int {

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	sort.Ints(nums)

	result := math.MaxInt
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, h := i+1, len(nums)-1
		for l < h {
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			if h < len(nums)-1 && nums[h] == nums[h+1] {
				h--
				continue
			}
			sum := nums[i] + nums[l] + nums[h]
			if abs(sum-target) < abs(result-target) {
				result = sum
			}

			if sum > target {
				h--
			} else if sum < target {
				l++
			} else {
				return sum
			}
		}
	}

	return result
}
