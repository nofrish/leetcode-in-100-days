package _033

func search_(nums []int, target int) int {

	binarySearch := func(nums []int, target int) int {

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

	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if idx := binarySearch(nums[:low], target); idx != -1 {
		return idx
	}
	if idx := binarySearch(nums[low:], target); idx != -1 {
		return low + idx
	}
	return -1
}
