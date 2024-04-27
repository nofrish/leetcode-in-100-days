package _283

func smallestDivisor(nums []int, threshold int) int {

	low, high := 1, 1
	for _, num := range nums {
		high = max(high, num)
	}

	minDivider := high
	for low <= high {
		mid := (low + high) >> 1
		if f(nums, mid) > threshold {
			low = mid + 1
		} else {
			minDivider = mid
			high = mid - 1
		}
	}
	return minDivider
}

// f 是单调递减函数，divider 越小，结果越大
func f(nums []int, divider int) int {
	sum := 0
	for _, num := range nums {
		sum += (num + divider - 1) / divider
	}
	return sum
}
