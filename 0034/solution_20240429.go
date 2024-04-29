package _034

func searchRange(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	start, end := 0, 0

	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) >> 1
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if nums[high] != target {
		return []int{-1, -1}
	}
	start = high

	low, high = 0, len(nums)-1
	for low < high {
		mid := (low + high + 1) >> 1
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid
		}
	}
	end = low

	return []int{start, end}
}
