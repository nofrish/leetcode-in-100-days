package _035

func searchInsert_(nums []int, target int) int {

	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > target {
			hi = mid - 1
		} else if nums[mid] < target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid + 1
			}
			lo = mid + 1
		} else {
			return mid
		}
	}
	return 0
}
