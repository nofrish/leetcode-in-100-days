package _033

func search(nums []int, target int) int {

	binarySearch := func(low, high int) int {
		for low <= high {
			mid := low + (high-low)/2
			if nums[mid] < target {
				low = mid + 1
			} else if nums[mid] > target {
				high = mid - 1
			} else {
				return mid
			}
		}
		return -1
	}

	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if idx := binarySearch(0, high-1); idx != -1 {
		return idx
	}
	if idx := binarySearch(high, len(nums)-1); idx != -1 {
		return idx
	}
	return -1
}
