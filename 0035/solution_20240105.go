package _035

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			high = mid - 1
		}
	}

	return len(nums)
}
