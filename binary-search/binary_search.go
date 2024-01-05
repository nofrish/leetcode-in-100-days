package binary_search

// BinarySearchBasic find any one number in nums EQUAL to target
func BinarySearchBasic(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// BinarySearchBasicByChatGPT is written by ChatGPT
func BinarySearchBasicByChatGPT(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		md := lo + (hi-lo)>>1
		switch {
		case nums[md] < target:
			lo = md + 1
		case nums[md] > target:
			hi = md - 1
		default:
			return md
		}
	}
	return -1
}

// BinarySearchFindFirst find the first number in nums EQUAL to target
func BinarySearchFindFirst(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// BinarySearchFindLast find the last number in nums EQUAL to target
func BinarySearchFindLast(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			low = mid + 1
		} else if nums[mid] < target {
			high = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			low = mid + 1
		}
	}

	return -1
}

// BinarySearchFindFirstGreaterOrEqual find the first number in nums GREATER or EQUAL to target
func BinarySearchFindFirstGreaterOrEqual(nums []int, target int) int {
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
	return -1
}

// BinarySearchFindLastLessOrEqual find the last number in nums LESS or EQUAL to target
func BinarySearchFindLastLessOrEqual(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}
