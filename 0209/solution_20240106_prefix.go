package _209

import "math"

func minSubArrayLen__(target int, nums []int) int {

	prefix := make([]int, len(nums)+1)
	for i := range nums {
		prefix[i+1] = prefix[i] + nums[i]
	}

	min := math.MaxInt
	for i, v := range prefix {
		nt := v + target
		low, high := i, len(prefix)-1
		for low <= high {
			mid := low + (high-low)/2
			if prefix[mid] < nt {
				low = mid + 1
			} else {
				if mid == i || prefix[mid-1] < nt {
					if min > mid-i {
						min = mid - i
					}
				}
				high = mid - 1
			}
		}
	}

	if min == math.MaxInt {
		min = 0
	}
	return min
}
