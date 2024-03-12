package _015

import "sort"

func threeSum___(nums []int) (result [][]int) {

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			}
		}
	}
	return
}
