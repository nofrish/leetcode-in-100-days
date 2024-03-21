package _015

import "sort"

func threeSum____(nums []int) (result [][]int) {

	sort.Ints(nums)

	for i, a := range nums {
		if a > 0 {
			return
		}
		if i > 0 && a == nums[i-1] {
			continue
		}

		low, high := i+1, len(nums)-1
		for low < high {
			if low > i+1 && nums[low] == nums[low-1] {
				low++
				continue
			}
			if high < len(nums)-1 && nums[high] == nums[high+1] {
				high--
				continue
			}
			sum := a + nums[low] + nums[high]
			if sum > 0 {
				high--
			} else if sum < 0 {
				low++
			} else {
				result = append(result, []int{a, nums[low], nums[high]})
				low++
				high--
			}
		}
	}
	return
}
