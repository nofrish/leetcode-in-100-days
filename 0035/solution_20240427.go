package _035

func searchInsert__(nums []int, target int) int {
	low, high := -1, len(nums)-1
	for low < high {
		mid := (low + high + 1) >> 1
		if nums[mid] < target {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low + 1
}
