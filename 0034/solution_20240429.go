package _034

func searchRange(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{-1, -1}
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
	if nums[lo] != target {
		return []int{-1, -1}
	}

	start, hi := lo, len(nums)-1
	for lo < hi {
		mid := (lo + hi + 1) >> 1
		if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid
		}
	}
	return []int{start, lo}
}
