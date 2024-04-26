package _081

func search(nums []int, target int) bool {

	binarySearch := func(low, high int) bool {
		for low <= high {
			mid := (low + high) >> 1
			if nums[mid] > target {
				high = mid - 1
			} else if nums[mid] < target {
				low = mid + 1
			} else {
				return true
			}
		}
		return false
	}

	low, high := 0, len(nums)-1
	for nums[low] == nums[high] && low < high {
		low++
	}

	for low < high {
		mid := (low + high) >> 1
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return binarySearch(0, low-1) || binarySearch(low, len(nums)-1)
}
