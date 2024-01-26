package _018

import "sort"

func fourSum(nums []int, target int) (result [][]int) {

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] > target && nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] > target && nums[j] > 0 {
				break
			}
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, len(nums)-1
			for left < right {
				if left > j+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				if right < len(nums)-1 && nums[right] == nums[right+1] {
					right--
					continue
				}
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
				}
			}
		}
	}

	return
}
