package _015

import "sort"

func threeSum_(nums []int) (result [][]int) {

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // deduplicate first
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
			if sum < 0 {
				left++
			} else if sum > 0 {
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
