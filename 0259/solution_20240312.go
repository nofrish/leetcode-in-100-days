package _259

import "sort"

func threeSumSmaller(nums []int, target int) (result int) {

	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)

	for i := range nums {
		low, high := i+1, len(nums)-1
		for low < high {
			sum := nums[i] + nums[low] + nums[high]
			if sum < target {
				result = result + (high - low)
				low++
			} else {
				high--
			}

		}
	}
	return
}
