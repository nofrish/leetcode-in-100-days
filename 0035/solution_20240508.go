package _035

func searchInsert___(nums []int, target int) int {

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return hi
}
