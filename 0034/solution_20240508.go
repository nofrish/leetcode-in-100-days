package _034

func searchRange_(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	lo, hi := 0, len(nums)-1
	for lo < hi { // use lo < hi to find the first element
		mid := (lo + hi) >> 1
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if nums[hi] != target { // check if we found first element at lo == hi
		return []int{-1, -1}
	}

	start, hi := lo, len(nums)-1
	for lo <= hi { // use lo <= hi to find the second element
		mid := (lo + hi) >> 1
		if nums[mid] > target {
			hi = mid - 1
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return []int{start, mid}
			}
			lo = mid + 1
		}
	}

	// should never happen
	return []int{-1, -1}
}
