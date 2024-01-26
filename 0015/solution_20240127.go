package _015

import "sort"

func threeSum_(nums []int) (result [][]int) {

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := 0 - nums[i]
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

			sum := nums[left] + nums[right]
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			}
		}
	}

	return
}
