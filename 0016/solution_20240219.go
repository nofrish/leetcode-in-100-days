package _016

import (
	"math"
	"sort"
)

func threeSumClosest_(nums []int, target int) (result int) {

	sort.Ints(nums)

	distance := math.MaxInt
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			sum := nums[i] + nums[left] + nums[right]
			if distance > int(math.Abs(float64(sum)-float64(target))) {
				result = sum
				distance = int(math.Abs(float64(sum) - float64(target)))
			}
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return target
			}
		}
	}
	return
}
