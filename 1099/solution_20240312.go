package _099

import (
	"sort"
)

func twoSumLessThanK(nums []int, k int) int {

	sort.Ints(nums)

	low, high, result := 0, len(nums)-1, -1
	for low < high {
		sum := nums[low] + nums[high]
		if sum >= k {
			high--
		} else {
			result = max(result, sum)
			low++
		}
	}
	return result
}
