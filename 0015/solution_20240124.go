package _015

import "sort"

func threeSum(nums []int) (result [][]int) {

	sort.Ints(nums)
	for i, v := range nums {
		if v > 0 {
			return
		}
		if i > 0 && v == nums[i-1] {
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
			sum := v + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				result = append(result, []int{v, nums[left], nums[right]})
				left++
				right--
			}
		}
	}
	return
}
